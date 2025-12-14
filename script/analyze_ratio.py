#!/usr/bin/env python3
"""
分析 logprocess 日志文件，统计 EndRound 事件中 ratio >= 2.0 的情况
"""
import json
import sys
from datetime import datetime

def analyze_log(log_file):
    """分析日志文件中的 EndRound 事件"""

    total_end_rounds = 0
    high_ratio_count = 0
    high_ratio_details = []

    print(f"Analyzing {log_file}...")
    print("=" * 100)

    try:
        with open(log_file, 'r', encoding='utf-8') as f:
            for line in f:
                line = line.strip()
                if not line:
                    continue

                try:
                    log_entry = json.loads(line)

                    # 只处理 msg=RoundState Updated 且 action=EndRound 的日志
                    if log_entry.get('msg') == 'RoundState Updated' and log_entry.get('action') == 'EndRound':
                        total_end_rounds += 1

                        epoch = log_entry.get('epoch')
                        ratio = log_entry.get('ratio', 0)
                        minority_side = log_entry.get('minoritySide', '')
                        bull_amount = log_entry.get('bullAmount', '0')
                        bear_amount = log_entry.get('bearAmount', '0')
                        ts = log_entry.get('ts', '')

                        # 统计 ratio >= 2.0
                        if ratio >= 2.0:
                            high_ratio_count += 1
                            high_ratio_details.append({
                                'epoch': epoch,
                                'ratio': ratio,
                                'minority_side': minority_side,
                                'bull_amount': bull_amount,
                                'bear_amount': bear_amount,
                                'timestamp': ts,
                            })

                except json.JSONDecodeError:
                    continue

    except FileNotFoundError:
        print(f"Error: File '{log_file}' not found")
        return
    except Exception as e:
        print(f"Error: {e}")
        return

    # 打印统计结果
    print(f"\nSummary:")
    print("-" * 100)
    print(f"Total EndRound events: {total_end_rounds}")
    print(f"High ratio rounds (ratio >= 2.0): {high_ratio_count}")

    if total_end_rounds > 0:
        percentage = (high_ratio_count / total_end_rounds) * 100
        print(f"Percentage: {percentage:.2f}%")

    # 打印详情
    if high_ratio_details:
        print(f"\nHigh Ratio Details (ratio >= 2.0):")
        print("-" * 100)
        print(f"{'Epoch':<10} {'Time':<22} {'Ratio':<10} {'Minority':<10} {'Bull Amount':<25} {'Bear Amount':<25}")
        print("-" * 100)

        for item in high_ratio_details:
            print(f"{item['epoch']:<10} {item['timestamp']:<22} {item['ratio']:<10.4f} {item['minority_side']:<10} {item['bull_amount']:<25} {item['bear_amount']:<25}")

    else:
        print(f"\nNo high ratio opportunities (ratio >= 2.0) found.")

    print("=" * 100)

if __name__ == '__main__':
    if len(sys.argv) > 1:
        log_file = sys.argv[1]
    else:
        # 默认使用今天的日志文件
        today = datetime.now().strftime('%Y-%m-%d')
        log_file = f'logprocess_{today}.log'

    analyze_log(log_file)
