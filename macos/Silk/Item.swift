//
//  Item.swift
//  Silk
//
//  Created by Mark Steve Samson on 5/3/22.
//

import Foundation

struct Item: Hashable, Codable {
    var name: String
    var data: Data
    var ts: Date
}

