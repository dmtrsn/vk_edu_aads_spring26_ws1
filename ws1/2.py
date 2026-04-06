def stack_sort(stack):
    tmp= []
    
    while stack:
        tmp1 = stack.pop()

        while tmp and tmp[-1] < tmp1:
            stack.append(tmp.pop())

        tmp.append(tmp1)

    return tmp

stack = [7, 1, 3, 4]
sorted_stack = stack_sort(stack)
print(sorted_stack)