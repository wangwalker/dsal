//
//  Queue.swift
//  Algorithms
//
//  Created by Walker Wang on 2021/11/17.
//

import UIKit

struct Queue<Element: CustomStringConvertible> {
    private var items: [Element] = []
    public var size: Int {
        return items.count
    }
    
    mutating func enqueue(_ item: Element) {
        items.append(item)
    }
    mutating func dequeue() -> Element? {
        if items.isEmpty {
            return nil
        }
        let item = items.first
        items.remove(at: 0)
        return item
    }
}

extension Queue: CustomStringConvertible {
    var description: String {
       "[" + items.map { $0.description } .joined(separator: ", ") + "]"
    }
}
