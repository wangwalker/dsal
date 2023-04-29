//
//  Tree.swift
//  Algorithms
//
//  Created by Walker Wang on 2021/11/20.
//

import Foundation

fileprivate protocol Nodable {
    associatedtype Key: Comparable
    associatedtype Value: Comparable
    var key: Key { get }
    var value: Value { set get }
    var count: Int { set get }
}

class TreeNode: Nodable {
    var key: Int
    var value: Int
    var count: Int = 0
    var left: TreeNode?
    var right: TreeNode?
    init(k: Int, v: Int) {
        key = k
        value = v
    }
    init(k: Int, v: Int, c: Int) {
        key = k
        value = v
        count = c
    }
}

class Tree {
    var root: TreeNode
    init(r: TreeNode) {
        root = r
    }
    func size() -> Int {
        return size(node: root)
    }
    private func size(node: TreeNode?) -> Int {
        if node == nil { return 0 }
        else { return node!.count }
    }
}
