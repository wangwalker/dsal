//
//  Sort.swift
//  Algorithms
//
//  Created by Walker Wang on 2021/11/20.
//

import UIKit

protocol Sortable {
    associatedtype Element: Comparable
    var items: [Element] { get }
    var swappingCount: Int { get }
    mutating func sort()
    mutating func swap(_ i: Int, _ j: Int)
    func log()
}

class Sort: Sortable {
    var items: [Int]
    var swappingCount: Int = 0
    
    init(items: [Int]) {
        self.items = items
    }
    func sort() {
        print("override...")
    }
    func swap(_ i: Int, _ j: Int) {
        items.swapAt(i, j)
        swappingCount += 1
    }
    func log() {
        print("Sort log: swapping count = \(swappingCount)")
    }
}
extension Sort {
    public var description: String {
        get {
            "[" + items.map { $0.description } .joined(separator: ", ") + "]"
        }
    }
}

class SelectionSort: Sort {
    override func sort() {
        for outer in 0..<items.count {
            var minIndex = outer
            var min = items[outer]
            
            for inner in outer+1..<items.count {
                if min > items[inner] {
                    minIndex = inner
                    min = items[inner]
                }
            }
            if outer != minIndex {
                swap(outer, minIndex)
            }
        }
        log()
    }
}

class InsertionSort: Sort {
    override func sort() {
        for outer in 1..<super.items.count {
            let current = items[outer];
            for inner in 0..<outer {
                if items[inner] > current {
                    swap(outer, inner)
                }
            }
        }
        log()
    }
}

/// 泛化的插入排序
/// 在插入排序中，比较的幅度为1；在希尔排序中，比较的幅度从大概n/3，一直降到1
class ShellSort: Sort {
    override func sort() {
        let n: Int = items.count
        var step: Int = 1
        while step <= n/3 {
            step = 3*step + 1
        }
        while step >= 1 {
            for i in step..<n {
                for j in stride(from: i, through: 0, by: -step) {
                    if j >= step && items[j] < items[j-step] {
                        swap(j, j-step)
                    }
                }
            }
            step = step/3
        }
        log()
    }
}

/// 快速排序
/// 快速排序也是分而治之的思想，每次至少能将一个元素放在合适的位置，时间复杂度是nlog(n)
/// 快速排序虽然很快，但非常脆弱，如果元素很少或者有序性较好，可能退回平方级别
class QuickSort: Sort {
    override func sort() {
        sort(lo: 0, hi: items.count-1)
        log()
    }
    private func sort(lo: Int, hi: Int) {
        if hi <= lo {
            return
        }
        let j = partion(lo: lo, hi: hi)
        sort(lo: lo, hi: j-1)
        sort(lo: j+1, hi: hi)
    }
    private func partion(lo: Int, hi: Int) -> Int {
        var i = lo
        var j = hi + 1  // 第一个元素作为切分点，后面的比较过程会将其越过，但最后一个元素不能被越过，所以j向后挪了一位
        let v = items[i]
        while true {
            /// left <= v
            i += 1
            j -= 1
            while i < items.count && items[i] <= v {
                if i == hi {
                    break
                }
                i += 1
            }
            /// v <= right
            while v <= items[j] {
                if j == lo {
                    break
                }
                j -= 1
            }
            if i >= j {
                break
            }
            swap(i, j)
        }
        swap(lo, j)
        return j
    }
}

/// 归并排序
/// 其思想也是分而治之，先将数组分成两半，分别排好序，然后再将这两半合并在一起
/// 除此之外，还需要额外的存储空间
class MergeSort: Sort {
    private var aux: [Int] = []
    private var recursiveCount = 0
    override func sort() {
        aux = Array.init(repeating: 0, count: items.count)
        sort(lo: 0, hi: items.count-1)
        print("recursive count is \(recursiveCount)")
    }
    private func sort(lo: Int, hi: Int) {
//        if lo >= hi {
//            return
//        }
        /// 当子数组长度小于3时，直接用插入排序
        if hi - lo <= 3 {
            /// 至少有两个元素
            if hi - lo >= 1  {
                for i in (lo+1)...hi {
                    for j in lo..<i {
                        if items[j] > items[i] {
                            swap(i, j)
                        }
                    }
                }
            }
            return
        }
        recursiveCount += 1
        let mid = (lo + hi) / 2
        sort(lo: lo, hi: mid)
        sort(lo: mid+1, hi: hi)
        merge(lo: lo, mi: mid, hi: hi)
    }
    private func merge(lo: Int, mi: Int, hi: Int) {
        var i = lo      /// 左半部分的开始
        var j = mi + 1  /// 右半部分的开始
        for k in lo...hi {
            aux[k] = items[k]
        }
        for k in lo...hi {
            /// 如果左边用完，取右边的部分
            if i > mi {
                items[k] = aux[j]
                j += 1
            }
            /// 如果右半部分用完，取左半部分的
            else if j > hi {
                items[k] = aux[i]
                i += 1
            }
            /// 如果右半部分小于左半部分，取右半部分的
            else if aux[j] < aux[i] {
                items[k] = aux[j]
                j += 1
            }
            /// 否则，左半部分更小，取左半部分的
            else {
                items[k] = aux[i]
                i += 1
            }
        }
    }
}
