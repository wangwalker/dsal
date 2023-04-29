// swift-tools-version: 5.7
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "dsal",
    products: [
        // Products define the executables and libraries a package produces, and make them visible to other packages.
        .library(
            name: "Linear",
            targets: ["Linear"]),
        .library(
            name: "Graph",
            targets: ["Graph"]),
        .library(
            name: "Tree",
            targets: ["Tree"]),
        .library(
            name: "Sort",
            targets: ["Sort"]),
    ],
    dependencies: [
        // Dependencies declare other packages that this package depends on.
        // .package(url: /* package url */, from: "1.0.0"),
    ],
    targets: [
        // Targets are the basic building blocks of a package. A target can define a module or a test suite.
        // Targets can depend on other targets in this package, and on products in packages this package depends on.
        .target(
            name: "Linear",
            dependencies: []),
        .target(
            name: "Graph",
            dependencies: []),
        .target(
            name: "Tree",
            dependencies: []),
        .target(
            name: "Sort",
            dependencies: []),
        .testTarget(
            name: "dsalTests",
            dependencies: []),
    ]
)
