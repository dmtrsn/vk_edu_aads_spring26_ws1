class Node:
    def __init__(self, val, next=None):
        self.value = val
        self.nextNode = next


def getIntersectionNode(h1, h2):
    if not h1 or not h2:
        return None
    
    p1, p2 = h1, h2

    while p1 != p2:
        p1 = p1.nextNode if p1 else h2
        p2 = p2.nextNode if p2 else h1

    return p2


common = Node(8, Node(4, Node(5)))

headA = Node(4, Node(1, common))

headB = Node(5, Node(6, Node(1, common)))

intersection = getIntersectionNode(headA, headB)

if intersection:
    print("Intersection at node with value:", intersection.value)
else:
    print("No intersection")