# Solving the area problem with shapely

from shapely import Polygon

x = y = border = 0
points = [(0,0)]
#f = 'sample.txt'
f = 'input.txt'
for l in open(f):
    
    d, n, c = l.split()
    n = int(n)
    if d == 'R': x += n
    if d == 'L': x -= n
    if d == 'U': y -= n
    if d == 'D': y += n
    points.append((x,y))
    border += n + .5

points.append((0,0))
poly = Polygon(points)
print(poly.area + border/2 - 2.5)
