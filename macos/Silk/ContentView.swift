//
//  ContentView.swift
//  Silk
//
//  Created by Mark Steve Samson on 5/3/22.
//

import SwiftUI

struct ContentView: View {
    @State private var linkCode: String = ""
    var items: [Item]
    var options: Options?
    var onLink: ((String) -> Void)?
    var onSync: (() -> Void)? = nil
    var body: some View {
        List {
            Text("Silk").font(.title)
            Button("Sync", action: {
                onSync?()
            })
            ForEach(items, id: \.name) { item in
                Text(item.name).padding()
            }
            Text("DB Directory: \(options!.Dir)")
            TextField("Link Code", text: $linkCode)    .onSubmit {
                    onLink?(linkCode)
                }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView(items: [
            Item(name: "test1", data: Data(), ts: Date()),
            Item(name: "test2", data: Data(), ts: Date()),
            Item(name: "test3", data: Data(), ts: Date()),
        ])
    }
}

