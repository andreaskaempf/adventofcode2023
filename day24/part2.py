#!/usr/bin/env python3
#
# Z3 constaint programming solution for Part 2.

import z3

# Parse input file into lists of numbers, 6 for each stone
# e.g., 19, 13, 30 @ -2,  1, -2
stones = []
for line in open('input.txt'):
    words = line.split()
    words = words[:3] + words[4:]
    words = [w.strip().replace(',', '') for w in words]
    nums = [int(w) for w in words]
    stones.append(nums)
print(len(stones), 'stones read')

# Create a solver
print('Creating solver')
solver = z3.Solver()

# Create six variables, for the position and velocity of the new stone
x = z3.BitVec('x', 64)
y = z3.BitVec('y', 64)
z = z3.BitVec('z', 64)
vx = z3.BitVec('vx', 64)
vy = z3.BitVec('vy', 64)
vz = z3.BitVec('vz', 64)

# Add constraints
# - The position of a stone is always its initial position plus its velocity
#   times time, i.e., x + vx * t == ax + vax * t    
# - For each existing stone, at some time t, the position of the stone must be
#   the same as that of the new stone
# - The time t must be positive
i = 0
for s in stones: #[:10]:

    # Unpack the variables for this stone
    x1, y1, z1, vx1, vy1, vz1 = s

    # Create a time variable for each stone, since each stone will
    # intersect the trajectory of the new stone at a different time
    i += 1
    t = z3.BitVec('t%d' % i, 64)

    # Add constraints, that the current stone must have the same position as
    # the new stone at some time t, one constraint per dimension
    solver.add(t >= 0)  # time must be positive
    solver.add(x + vx * t == x1 + vx1 * t)
    solver.add(y + vy * t == y1 + vy1 * t)
    solver.add(z + vz * t == z1 + vz1 * t)

# Check the model, find the solution
print('Solving')
assert solver.check() == z3.sat

# Get the model inputs and show answer
model = solver.model()
x = model.eval(x).as_long()
y = model.eval(y).as_long()
z = model.eval(z).as_long()
print('x/y/z:', x, y, z)
print('Part 2:', x + y + z)
