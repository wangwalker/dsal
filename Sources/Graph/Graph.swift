//
//  Graph.swift
//  Algorithms
//
//  Created by Walker Wang on 2021/11/17.
//

import Foundation

fileprivate protocol Nodeable {
    var index: Int { set get }
    var name: String { set get }
    var visited: Bool { get }
}

fileprivate protocol Visiting {
    mutating func visit(callback: () -> Void)
}

struct GraphNode: Nodeable, Visiting {
    var index: Int
    var name: String
    var visited: Bool = false
    init(index: Int, name: String) {
        self.index = index
        self.name = name
    }
    mutating func visit(callback: () -> Void) {
        self.visited = true
        callback()
    }
}

/**
 * 图的抽象数据类型
 */
protocol Graph {
    associatedtype Node: Comparable
    /// 顶点数量
    var vertices: Int { get }
    /// 边的数量
    var edges: Int { get }
    /// 在顶点v和w之间添加一条边
    mutating func addEdge(v: Node, w: Node, wt: Double)
    /// 返回顶点v的所有相邻顶点
    func adjacents(v: Node) -> Array<Node>
}

/**
 * 图的搜索算法，包括深度优先和广度优先
 */
protocol Searchable {
    func search(s: Int)
    /// 顶点v和s是连通的吗
    func marked(v: Int) -> Bool
    /// 从s到v之前的最长路径
    func path(to v: Int) -> [Int]
}

/**
 * 图的连通性测试
 * 连通图是指，从图的任意顶点开始，都存在一条路径，可以到达另一个任意顶点
 * 一个非连通图，是由若干连通的部分组成的，它们都是极大连通子图，也称为连通分量
 */
protocol Connectable {
    /// 如果vertices中的所有顶点都处于同一个连通分量，则它们是连通的
    func connected(vertices: [Int]) -> Bool
}

/**
 * 检测有向图中是否含有环
 * 如果含有环，从某个顶点开始，存在一条路径，能够再次回到该顶点
 */
protocol Cyclable {
    /// 从起点s开始，如果存在一个环，返回环上的所有顶点
    func cycle(start: Int) -> [Int]
    func hasCycle() -> Bool
}

/**
 * 简单无向图。
 * 用“邻接表”作为图的表示方式，顶点用编号0~(v-1)标识
 */
class BasicGraph: Graph {
    public var vertices: Int = 0
    public var edges: Int = 0
    private var adjacents: [[Int]] = []
    
    init(v: Int) {
        self.vertices = v
        self.adjacents = Array.init(repeating: [], count: v)
    }
    init(v: Int, adj: [[Int]]) {
        self.vertices = v
        self.adjacents = adj
        self.edges = adj.compactMap({ vadj in
            return vadj.count
        }).reduce(0 , { partialResult, e in
            return partialResult + e
        })/2
    }
    func addEdge(v: Int, w: Int, wt: Double) {
        guard v < self.vertices, w < self.vertices else {
            return
        }
        self.adjacents[v].append(w)
        self.adjacents[w].append(v)
        self.edges += 1
    }
    
    func adjacents(v: Int) -> Array<Int> {
        if v < 0 ||  v >= self.vertices {
            return []
        }
        return self.adjacents[v]
    }
}

/**
 * 简单有向图（无权重）
 * 用“邻接表”作为图的表示方式，顶点用编号0~(v-1)标识
 */
class DirectedGraph: Graph {
    public var vertices: Int = 0
    public var edges: Int = 0
    private var adjacents: [[Int]] = []
    
    init(v: Int) {
        self.vertices = v
        self.adjacents = Array.init(repeating: [], count: v)
    }
    init(v: Int, adj: [[Int]]) {
        self.vertices = v
        self.adjacents = adj
        self.edges = adj.compactMap({ vadj in
            return vadj.count
        }).reduce(0 , { partialResult, e in
            return partialResult + e
        })
    }
    func addEdge(v: Int, w: Int, wt: Double) {
        guard v < self.vertices, w < self.vertices else {
            return
        }
        self.adjacents[v].append(w)
        self.edges += 1
    }
    func adjacents(v: Int) -> Array<Int> {
        if v < 0 ||  v >= self.vertices {
            return []
        }
        return self.adjacents[v]
    }
    /// 取反，将所有边的方向取反，这在某些场景中非常有用
    func reversed() -> DirectedGraph {
        let redi = DirectedGraph.init(v: vertices)
        for v in 0..<vertices {
            for w in adjacents(v: v) {
                redi.addEdge(v: w, w: v, wt: 0)
            }
        }
        return redi
    }
}


/**
 * func dfs(graph: Graph, start: Node) {
 *    visit(start)
 *    for node in graph.adjacents(node: start) {
 *      if node is not visited {
 *       dfs(graph, node)
 *      }
 *    }
 *  }
 */
class DepthFirstSearch: Searchable {
    private var graph: BasicGraph
    private var start: Int = 0
    /// 标记每个顶点的前一个顶点
    private var priors: [Int] = []
    private var marked: [Bool] = []
    
    init(g: BasicGraph) {
        graph = g
        marked = Array.init(repeating: false, count: graph.vertices)
        priors = Array.init(repeating: 0, count: graph.vertices)
    }
    
    func search(s: Int) {
        start = s
        dfs(graph: graph, v: s)
    }
    func marked(v: Int) -> Bool {
        return marked[v]
    }
    func path(to v: Int) -> [Int] {
        if !marked(v: v) {
            return []
        }
        var paths: [Int] = []
        var w: Int = v
        while w != start {
            paths.append(w)
            w = priors[w]
        }
        paths.append(start)
        return paths
    }
    private func dfs(graph: BasicGraph, v: Int) {
        marked[v] = true
        for adj in graph.adjacents(v: v) {
            if !marked[adj] {
                priors[adj] = v
                dfs(graph: graph, v: adj)
            }
        }
    }
}

class ConnectedComponent: Connectable {
    private var graph: BasicGraph
    private var components: [Int] = []
    private var marked:[Bool] = []
    private var loops: Int = 0
    
    init(graph: BasicGraph) {
        self.graph = graph
        components = Array.init(0..<graph.vertices)
        marked = [Bool](repeating: false, count: graph.vertices)
        for v in 0..<graph.vertices {
            if !marked[v] {
                dfs(graph: graph, v: v)
            }
            loops += 1
        }
    }
    func connected(vertices: [Int]) -> Bool {
        if vertices.isEmpty {
            return false
        }
        let first = components[vertices.first!]
        for v in vertices {
            if components[v] != first {
                return false
            }
        }
        return true
    }
    private func dfs(graph: BasicGraph, v: Int) {
        marked[v] = true
        components[v] = loops
        for adj in graph.adjacents(v: v) {
            if !marked[adj] {
                dfs(graph: graph, v: adj)
            }
        }
    }
}

class CycleFinder: Cyclable {
    private var graph: DirectedGraph
    private var marked: [Bool] = []
    /// 表示是否在当前调用栈上，如果在，意味着已经成环
    private var onstack: [Bool] = []
    private var priors: [Int] = []
    private var cycle = Queue<Int>()
    
    init(g: DirectedGraph) {
        graph = g
        marked = [Bool](repeating: false, count: graph.vertices)
        onstack = [Bool](repeating: false, count: graph.vertices)
        priors = [Int](repeating: -1, count: graph.vertices)
        for i in 0..<graph.vertices {
            if !marked[i] {
                dfs(graph: graph, v: i)
            }
        }
    }
    
    func cycle(start: Int) -> [Int] {
        marked = [Bool](repeating: false, count: graph.vertices)
        onstack = [Bool](repeating: false, count: graph.vertices)
        priors = [Int](repeating: -1, count: graph.vertices)
        cycle = Queue<Int>()
        dfs(graph: graph, v: start)
        var paths: [Int] = []
        while cycle.size > 0 {
            paths.append(cycle.dequeue()!)
        }
        return paths
    }
    func hasCycle() -> Bool {
        return cycle.size > 0
    }
    private func dfs(graph: DirectedGraph, v: Int) {
        marked[v] = true
        /// 当前节点入栈
        onstack[v] = true
        for w in graph.adjacents(v: v) {
            if hasCycle() {
                return
            }
            else if !marked[w] {
                priors[w] = v
                dfs(graph: graph, v: w)
            }
            else if onstack[w] {
                var x = v
                while x != w {
                    cycle.enqueue(x)
                    x = priors[x]
                }
                cycle.enqueue(w)
                cycle.enqueue(v)
            }
        }
        /// 当前节点出栈
        onstack[v] = false
    }
}

/**
 * bfs(graph: Graph, start: Node) {
 *  var queue: Queue<Node> = Queue()
 *  queue.enqueue(start)
 *  visit(start)
 *  while queue is not empty {
 *   let node = queue.dequeue()
 *   for v in graph.adjacents(node: v) {
 *    if v is not visited {
 *     visit(v)
 *     queue.enqueue(v)
 *     }
 *    }
 *   }
 */
class BreadthFirstSearch: Searchable {
    private var graph: BasicGraph
    private var start: Int = 0
    private var vertices: [Int] = []
    private var marked: [Bool] = []
    
    init(g: BasicGraph) {
        graph = g
    }
    func search(s: Int) {
        start = s
        vertices = [Int](repeating: 0, count: graph.vertices)
        marked = [Bool](repeating: false, count: graph.vertices)
        bfs(graph: graph, v: s)
    }
    func path(to v: Int) -> [Int] {
        if !marked[v] {
            return []
        }
        var paths: [Int] = []
        var w: Int = v
        while w != start {
            paths.append(w)
            w = vertices[w]
        }
        paths.append(start)
        return paths
    }
    func marked(v: Int) -> Bool {
        return marked[v]
    }
    private func bfs(graph: BasicGraph, v: Int) {
        var queue = Queue<Int>()
        queue.enqueue(v)
        marked[v] = true
        while queue.size > 0 {
            let vertex: Int = queue.dequeue()!
            graph.adjacents(v: vertex).forEach { a in
                if !marked[a] {
                    marked[a] = true
                    vertices[a] = vertex
                    queue.enqueue(a)
                }
            }
        }
    }
}

/**
 * 加权有向图中的边的定义
 */
protocol DirectedEdging {
    associatedtype Node: Comparable
    var from: Node { get }
    var to: Node { get }
    var weight: Double { get }
}

struct DirectedEdge: DirectedEdging {
    var from: Int
    var to: Int
    var weight: Double
    
    init(from: Int, to: Int, w: Double) {
        self.from = from
        self.to = to
        self.weight = w
    }
}

class EdgeWeigtedDiGraph: Graph {
    typealias Node = Int
    var vertices: Int
    var edges: Int = 0
    private var adjacents: [[DirectedEdge]] = []
    
    init(v: Int) {
        vertices = v
        adjacents = [[DirectedEdge]](repeating: [], count: v)
    }
    init(v: Int, adj: [[DirectedEdge]]) {
        vertices = v
        adjacents = adj
        edges = adj.compactMap({ es in
            return es.count
        }).reduce(0, { partialResult, c in
            return partialResult + c
        })
    }
    func addEdge(v: Int, w: Int, wt: Double) {
        let edge = DirectedEdge.init(from: v, to: w, w: wt)
        adjacents[v].append(edge)
        edges += 1
    }
    func adjacents(v: Int) -> Array<Int> {
        return adjacents[v].map { de in
            return de.to
        }
    }
    func adjacentEdges(v: Int) -> [DirectedEdge] {
        return adjacents[v]
    }
}

class ShortestPath {
    private var graph: EdgeWeigtedDiGraph
    private var start: Int
    private var distance: [Double] = []
    private var priors: [Int] = []
    private var processed:  [Int: Bool] = [:]
    
    init(g: EdgeWeigtedDiGraph, s: Int) {
        graph = g
        start = s
        priors = [Int](repeating: -1, count: g.vertices)
        distance = [Double](repeating: Double(Int.max), count: g.vertices)
        /// 更新起点的邻接点的权重
        for e in graph.adjacentEdges(v: s) {
            distance[e.to] = e.weight
        }
        var next = shortest()
        while next != -1 {
            relex(v: next)
            next = shortest()
        }
    }
    func distance(to v: Int) -> Double {
        return distance[v]
    }
    func path(to v: Int) -> [Int] {
        var paths: [Int] = []
        var w: Int = v
        while w != start && w >= 0 {
            paths.append(w)
            w = priors[w]
        }
        paths.append(start)
        return paths
    }
    /// 边的松弛是指，如果[start->w->v]的路径权重和比[start->v]更短，就说从start经过w到v是松弛的
    /// 这意味着，在算最短路径时，要更新从start到v的权重记录到目前状态
    private func relex(e: DirectedEdge) {
        let v = e.from
        let w = e.to
        if distance[w] > distance[v] + e.weight {
            distance[w] = distance[v] + e.weight
            priors[w] = v
        }
    }
    /// 顶点的松弛和边的松弛类似，只不过顶点会计算与之相连的所有边
    private func relex(v: Int) {
        if processed.keys.contains(v) && processed[v]! {
            return
        }
        for e in graph.adjacentEdges(v: v) {
            relex(e: e)
        }
        processed[v] = true
    }
    private func shortest() -> Int {
        var shortestNode = -1
        var shortest = Double(Int.max)
        for v in 1..<graph.vertices {
            if distance[v] < shortest && !processed.keys.contains(v) {
                shortestNode = v
                shortest = distance[v]
            }
        }
        return shortestNode
    }
}
