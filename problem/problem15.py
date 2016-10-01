# python2.7

def count_through_pattern(arg):
    outside_map = []
    for line in range(0, arg+1):
        inside_map = []
        for row in range(0, arg+1):
            # line:1 or row:1 has only 1 pattern.
            if line == 0 or row == 0:
                inside_map.append(1)
            else:
                top = outside_map[line-1][row]
                left = inside_map[row-1]
                inside_map.append(top+left)
        outside_map.append(inside_map)
    return outside_map[arg][arg]

if __name__ == '__main__':
    for i in range(1, 21):
        ret = count_through_pattern(i)
        print("{0}: {1}".format(i, ret))
