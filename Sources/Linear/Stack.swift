//
//  Stack.swift
//  Algorithms
//
//  Created by Walker Wang on 2021/11/17.
//

import Foundation

struct Stack<Element: CustomStringConvertible> {
    private var items: [Element] = []
    public var top: Element? {
        return items.last
    }
    public var size: Int {
        return items.count
    }
    
    mutating func push(_ item: Element) {
        items.append(item)
    }
    mutating func pop() -> Element? {
        return items.popLast()
    }
}

extension Stack: CustomStringConvertible {
    var description: String {
       "[" + items.map { $0.description } .joined(separator: ", ") + "]"
    }
}
