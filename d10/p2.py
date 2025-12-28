from dataclasses import dataclass
from typing import List
from z3 import Int, Optimize, Sum, sat

# ring the shame bell


@dataclass
class Machine:
    state: List[bool]
    buttons: List[List[int]]
    reqs: List[int]


def parse_machines(data: List[str]) -> List[Machine]:
    machines: List[Machine] = []

    for d in data:
        state_end = d.index(']')
        state: List[bool] = [False] * (state_end - 1)

        for s in range(1, state_end):
            state[s - 1] = (d[s] == '#')

        buttons_end = d.index('{')
        buttons_section = d[state_end + 2: buttons_end - 1].strip()

        buttons: List[List[int]] = []
        if buttons_section:
            buttons_raw = buttons_section.split(" ")
            for button_raw in buttons_raw:

                sub = button_raw[1:-1]
                if sub.strip() == "":
                    buttons.append([])
                else:
                    buttons.append([int(x) for x in sub.split(",")])

        reqs_section = d[buttons_end + 1: -1].strip()
        reqs: List[int] = []
        if reqs_section:
            reqs = [int(x) for x in reqs_section.split(",")]

        machines.append(Machine(state=state, buttons=buttons, reqs=reqs))

    return machines


def p2(desired, buttons):
    n = len(desired)
    m = len(buttons)

    opt = Optimize()
    x = [Int(f"x_{j}") for j in range(m)]
    opt.add([xj >= 0 for xj in x])

    for i in range(n):
        column_sum = Sum([x[j] for j, b in enumerate(buttons) if i in b])
        opt.add(column_sum == desired[i])

    opt.minimize(Sum(x))

    if opt.check() == sat:
        model = opt.model()
        return sum(model[xj].as_long() for xj in x)
    return 0


if __name__ == "__main__":

    with open(f"./data/d10")as f:
        data = f.readlines()
    data = [x.rstrip() for x in data]
    machines = parse_machines(data)
    ans = 0
    for m in machines:
        ans += p2(m.reqs, m.buttons)
    print(ans)
