#!/usr/bin/env python3
"""
PancakeSwap Prediction 合约事件获取脚本
支持分批获取多天的历史事件数据
"""

import json
import requests
import time
from typing import List, Dict, Any
from datetime import datetime

# 配置
# BSC 公开节点列表（按推荐顺序）：
# 1. https://bsc-dataseed.binance.org (官方)
# 2. https://bsc-dataseed1.binance.org (官方备用)
# 3. https://bsc-dataseed2.binance.org (官方备用)
# 4. https://bsc.publicnode.com (社区节点)
# 5. https://bsc-rpc.publicnode.com (社区节点)
RPC_URL = "https://bsc-dataseed.binance.org"
OUTPUT_FILE = "pancake_events.json"

# 合约地址
CONTRACTS = [
    "0x18B2A687610328590Bc8F2e5fEdDe3b582A49cdA",  # PancakeSwap Prediction V2 (BNB)
    "0x48781a7d35f6137a9135Bbb984AF65fd6AB25618",
    "0x7451F994A8D510CBCB46cF57D50F31F188Ff58F5"
]


# 事件签名哈希 (Topic0)
EVENT_TOPICS = {
    "BetBull": "0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12",
    "BetBear": "0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d",
    "Claim": "0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7",
    "LockRound": "0x85b533c0fa284d94993934e2b570c1e9b3b7d0bdb3e0ce92e65c26fd46f481a2",
    "StartRound": "0x0bb59eceb12f1bdd2f0b3c6e68bc76f3c93d328d251b1fca62a51f62f28c90a4",
    "EndRound": "0x070f615e3c25ace7a92f3a2f441c8d41bbfdafc6d641e2c544be47c0cd870e91"
}

# BSC 平均出块时间约 3 秒，一天约 28800 个区块
BLOCKS_PER_DAY = 28800
# 每次查询的最大区块数（避免节点限制）
MAX_BLOCKS_PER_QUERY = 1000
# 请求间隔（秒）- 避免触发限流
REQUEST_DELAY = 0.5
# 失败重试次数（设为很大的数，确保最终成功）
MAX_RETRIES = 50
# 初始重试间隔（秒），会逐渐增加
RETRY_DELAY_BASE = 3
# 最大重试间隔（秒）
RETRY_DELAY_MAX = 30


class EventFetcher:
    def __init__(self, rpc_url: str, contracts: List[str], output_file: str):
        self.rpc_url = rpc_url
        self.contracts = contracts
        self.output_file = output_file
        self.session = requests.Session()

    def rpc_call(self, method: str, params: List[Any]) -> Dict:
        """发送 JSON-RPC 请求"""
        payload = {
            "jsonrpc": "2.0",
            "method": method,
            "params": params,
            "id": 1
        }
        response = self.session.post(self.rpc_url, json=payload)
        response.raise_for_status()
        result = response.json()

        if "error" in result:
            raise Exception(f"RPC Error: {result['error']}")

        return result.get("result")

    def get_latest_block(self) -> int:
        """获取最新区块号"""
        block_hex = self.rpc_call("eth_blockNumber", [])
        return int(block_hex, 16)

    def get_logs(self, from_block: int, to_block: int, contract: str, batch_info: str = "") -> List[Dict]:
        """获取指定区块范围的事件日志（带指数退避重试机制）"""
        topics_list = list(EVENT_TOPICS.values())

        params = [{
            "fromBlock": hex(from_block),
            "toBlock": hex(to_block),
            "address": contract,
            "topics": [topics_list]  # OR 查询所有关注的事件
        }]

        # 指数退避重试机制
        for attempt in range(MAX_RETRIES):
            try:
                logs = self.rpc_call("eth_getLogs", params)
                if attempt > 0:
                    print(f"      ✓ 重试成功！")
                return logs if logs else []
            except Exception as e:
                if attempt < MAX_RETRIES - 1:
                    # 指数退避：等待时间逐渐增加
                    wait_time = min(RETRY_DELAY_BASE * (2 ** attempt), RETRY_DELAY_MAX)
                    print(f"      ⚠ {batch_info} 失败 (尝试 {attempt + 1}/{MAX_RETRIES}): {e}")
                    print(f"      ⏳ 等待 {wait_time:.0f} 秒后重试...")
                    time.sleep(wait_time)
                else:
                    print(f"      ✗ {batch_info} 重试 {MAX_RETRIES} 次后仍失败，请检查节点")
                    raise e

    def identify_event(self, log: Dict) -> str:
        """识别事件类型"""
        if not log.get("topics"):
            return "Unknown"

        topic0 = log["topics"][0]
        for event_name, topic_hash in EVENT_TOPICS.items():
            if topic0.lower() == topic_hash.lower():
                return event_name

        return "Unknown"

    def load_existing_data(self) -> Dict:
        """加载已存在的数据"""
        try:
            with open(self.output_file, 'r', encoding='utf-8') as f:
                return json.load(f)
        except FileNotFoundError:
            return {
                "metadata": {
                    "contracts": self.contracts,
                    "last_update": None,
                    "total_events": 0,
                    "block_range": {"from": None, "to": None}
                },
                "events": []
            }

    def save_data(self, data: Dict):
        """保存数据到文件"""
        with open(self.output_file, 'w', encoding='utf-8') as f:
            json.dump(data, f, indent=2, ensure_ascii=False)

    def fetch_events(self, days: int):
        """获取指定天数的事件数据"""
        print(f"开始获取最近 {days} 天的事件数据...")
        print(f"RPC: {self.rpc_url}")
        print(f"合约数量: {len(self.contracts)}")
        print(f"监控事件: {', '.join(EVENT_TOPICS.keys())}")
        print("=" * 60)

        # 获取最新区块
        latest_block = self.get_latest_block()
        print(f"最新区块: {latest_block}")

        # 计算起始区块
        total_blocks = days * BLOCKS_PER_DAY
        start_block = latest_block - total_blocks
        print(f"查询范围: {start_block} -> {latest_block} (共 {total_blocks} 个区块)")
        print("=" * 60)

        # 加载已有数据
        data = self.load_existing_data()
        existing_event_count = len(data["events"])

        # 按天分批查询
        all_new_events = []
        current_block = start_block

        for day in range(days):
            day_start_block = current_block
            day_end_block = min(current_block + BLOCKS_PER_DAY - 1, latest_block)

            print(f"\n[第 {day + 1}/{days} 天] 查询区块 {day_start_block} -> {day_end_block}")

            day_events = []

            # 按合约循环
            for contract_idx, contract in enumerate(self.contracts):
                contract_events = []
                batch_current = day_start_block
                batch_count = 0
                batch_success = 0
                batch_failed = 0

                # 将一天的区块分批查询（每批最多 MAX_BLOCKS_PER_QUERY 个区块）
                while batch_current <= day_end_block:
                    batch_from = batch_current
                    batch_to = min(batch_current + MAX_BLOCKS_PER_QUERY - 1, day_end_block)
                    batch_count += 1

                    batch_info = f"合约 {contract[:10]}... 批次 {batch_count}/{(day_end_block - day_start_block + 1 + MAX_BLOCKS_PER_QUERY - 1) // MAX_BLOCKS_PER_QUERY} ({batch_from}-{batch_to})"

                    try:
                        logs = self.get_logs(batch_from, batch_to, contract, batch_info)

                        # 处理日志
                        for log in logs:
                            event_type = self.identify_event(log)
                            event_data = {
                                "contract": contract,
                                "event_type": event_type,
                                "block_number": int(log["blockNumber"], 16),
                                "transaction_hash": log["transactionHash"],
                                "log_index": int(log["logIndex"], 16),
                                "topics": log["topics"],
                                "data": log["data"],
                                "removed": log.get("removed", False)
                            }
                            contract_events.append(event_data)

                        batch_success += 1
                        # 显示进度（只在有多个批次时显示）
                        total_batches = (day_end_block - day_start_block + 1 + MAX_BLOCKS_PER_QUERY - 1) // MAX_BLOCKS_PER_QUERY
                        if total_batches > 1:
                            print(f"    ✓ {batch_info}: {len(logs)} 个事件")

                    except Exception as e:
                        batch_failed += 1
                        print(f"    ✗ {batch_info} 最终失败，程序将终止")
                        raise e

                    # 避免请求过快
                    time.sleep(REQUEST_DELAY)

                    batch_current = batch_to + 1

                day_events.extend(contract_events)
                success_rate = (batch_success / batch_count * 100) if batch_count > 0 else 0
                status_emoji = "✓" if batch_failed == 0 else "⚠"
                print(f"  {status_emoji} 合约 {contract_idx + 1}/{len(self.contracts)} ({contract[:10]}...): {len(contract_events)} 个事件 (成功 {batch_success}/{batch_count} 批, {success_rate:.1f}%)")

            all_new_events.extend(day_events)
            print(f"  ✓ 第 {day + 1} 天共获取 {len(day_events)} 个事件")

            # 每天保存一次
            data["events"].extend(day_events)
            data["metadata"]["last_update"] = datetime.now().isoformat()
            data["metadata"]["total_events"] = len(data["events"])
            data["metadata"]["block_range"]["from"] = start_block
            data["metadata"]["block_range"]["to"] = day_end_block
            self.save_data(data)
            print(f"  💾 已保存到 {self.output_file} (累计: {len(data['events'])} 个事件)")

            current_block = day_end_block + 1

            # 避免请求过快
            time.sleep(0.5)

        # 统计
        print("\n" + "=" * 60)
        print("✅ 数据获取完成！")
        print(f"本次新增事件: {len(all_new_events)} 个")
        print(f"总事件数: {len(data['events'])} 个 (之前: {existing_event_count})")

        # 按事件类型统计
        event_type_counts = {}
        for event in data["events"]:
            event_type = event["event_type"]
            event_type_counts[event_type] = event_type_counts.get(event_type, 0) + 1

        print("\n事件类型统计:")
        for event_type, count in sorted(event_type_counts.items()):
            print(f"  {event_type}: {count} 个")

        print(f"\n数据已保存至: {self.output_file}")


def main():
    """主函数"""
    print("=" * 60)
    print("PancakeSwap Prediction 合约事件获取工具")
    print("=" * 60)

    # 获取用户输入
    try:
        days = int(input("\n请输入要获取的天数 (例如: 7, 30, 90): ").strip())
        if days <= 0:
            print("❌ 天数必须大于 0")
            return
    except ValueError:
        print("❌ 请输入有效的数字")
        return

    # 创建获取器并执行
    fetcher = EventFetcher(RPC_URL, CONTRACTS, OUTPUT_FILE)

    try:
        fetcher.fetch_events(days)
    except KeyboardInterrupt:
        print("\n\n⚠ 用户中断，数据已保存至最后一次更新")
    except Exception as e:
        print(f"\n❌ 发生错误: {e}")
        import traceback
        traceback.print_exc()


if __name__ == "__main__":
    main()
