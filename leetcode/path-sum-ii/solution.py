def path_sum(node, total, parents=None):
    if parents is None:
        parents = []
    if node is None:
        return []
    parents.append(node.val)
    rtn = []
    if node.left is None and node.right is None:
        if total == sum(parents):
            return [parents]
    lpaths = rpaths = []
    if node.left:
        lpaths = path_sum(node.left, total, parents[:])
    if node.right:
        rpaths = path_sum(node.right, total, parents[:])
    return rtn + lpaths + rpaths
