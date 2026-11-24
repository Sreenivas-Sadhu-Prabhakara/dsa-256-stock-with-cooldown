# Best Time to Buy and Sell Stock With Cooldown

> NeetCode 150 · 2-D DP — a three-state machine (holding / cooling / free) with the transitions drawn straight from the rules; the DP is the automaton run forward. Go, stdlib only, tested.

## Problem

Trade one share repeatedly for maximum profit, but each sale forces a one-day cooldown before the next buy. The cooldown kills the simple "sum every rise" greedy.

## Approach

When rules constrain *sequences* of actions, model the day as a state machine: **hold** (own the share), **sold** (sold today, cooling), **rest** (flat, free to buy). Each rule becomes one transition — sell: `hold + price → sold`; buy (only from free): `rest − price → hold`; cooldown expiring: `sold → rest`. Run the automaton over the prices keeping the best profit per state; the answer is the better flat state at the end. Seeding `hold` with −∞ forbids selling before any buy. The same state-machine template handles transaction fees and k-transaction variants — only the states and arrows change.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
