//
//  SilkApp.swift
//  Silk
//
//  Created by Mark Steve Samson on 5/3/22.
//

import SwiftUI
import Silk_kv
import os

let kv = "silk"
let logger = Logger()

@main
struct SilkApp: App {
    @State private var items: [Item] = getItems()
    var body: some Scene {
        WindowGroup {
            ContentView(
                items: items,
                options: getOptions(),
                onLink: link,
                onSync: {
                    items = getItems()
                }
            )
        }
    }
}

struct Options: Decodable {
    var Dir: String
}

func getOptions() -> Options {
    let JSON = Silk_kvGetOptions(kv, nil)
    let jsonData = JSON.data(using: .utf8)!
    let options: Options = try! JSONDecoder().decode(Options.self, from: jsonData)
    return options
}

func getItems() -> [Item] {
    let JSON = Silk_kv.Silk_kvKeys(kv, nil)
    if JSON == "null" {
        return []
    }
    let jsonData = JSON.data(using: .utf8)!
    let keys: [String] = try! JSONDecoder().decode([String].self, from: jsonData)
    return keys.map({key in
        return Item(name: key, data: Data(), ts: Date())
    })
}

func link(code: String) {
    Silk_kvLink(code, nil)
    logger.log("Linked to \(code)!")
}

