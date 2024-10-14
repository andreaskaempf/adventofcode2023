// Advent of Code 2023, Day 05
//
// Given a list of "seeds" (numbers), go through a series of transformations
// based on tables. Each table is a tuple (dst, src, n). If the number is
// within src .. src+n,  it is assigned to dst. If no dst is found, the dst
// is the same as src for that transformation. For Part 1, find the lowest
// final transformation. Part 2 is the same, but treat each pair of "seeds"
// as start and length of a range. Brute force approach, acceptably fast
// in Go, Rust, and Zig.
//
// To build:
//   zig build                          (debug mode, runs in ~15 mins)
//   zig build -Doptimize=ReleaseFast   (fast mode, takes ~2 mins)
//
// To run:
//   zig-out/bin/day05
//
// AK, 13 Oct 2024 for Zig version

const std = @import("std");

// One section: name, list of ranges
const Section = struct {
    name: []const u8,
    ranges: std.ArrayList(Range),
};

// A range is just 3 numbers
const Range = struct {
    a: i64,
    b: i64,
    c: i64,
};

pub fn main() !void {

    // For writing output
    const stdout = std.io.getStdOut().writer();

    //  Get a memory allocator
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    // Read data
    const filename = "../input.txt";
    const d = try read_input(filename, allocator);
    const seeds = d.seeds;
    var sections = d.sections;

    // Show the data
    try stdout.print("Seeds: {any}\n", .{seeds.items});
    try stdout.print("Sections ({d}):\n", .{sections.items.len});
    for (sections.items) |s| {
        try stdout.print("  {s}: {d} ranges\n", .{ s.name, s.ranges.items.len });
    }

    // Part 1: process each seed, find the lowest location
    var lowest: i64 = -1;
    for (seeds.items) |s| {
        const loc = process_seed(s, sections);
        if (lowest == -1 or loc < lowest) {
            lowest = loc;
        }
    }
    try stdout.print("Part 1 (s/b 424490994): {d}\n", .{lowest});

    // Part 2: same, but treat seed numbers as ranges
    try stdout.print("Part 2:\n", .{});
    lowest = -1;
    var i: usize = 0;
    while (i < seeds.items.len) { // TODO: for loop with range, increment by 2 each time
        const s1 = seeds.items[i]; // start of range
        const n = seeds.items[i + 1]; // length of range
        try stdout.print("  seed {d} for length {d}\n", .{ s1, n });
        var j: i64 = 0; // TODO: can we use a 0..n for loop here (causes usize -> int error)
        while (j < n) {
            const loc = process_seed(s1 + j, sections);
            if (lowest == -1 or loc < lowest) {
                lowest = loc;
            }
            j += 1;
        }
        i += 2;
    }
    try stdout.print("Part 2 (s/b 15290096): {d}\n", .{lowest});

    // Free the sections and seeds data
    while (sections.items.len > 0) { // each section
        const sec = sections.pop();
        allocator.free(sec.name); // free name string
        sec.ranges.deinit(); // free up the list of ranges
    }
    seeds.deinit();
    sections.deinit();
}

// Process one seed number, by passing it through each section
fn process_seed(seed: i64, sections: std.ArrayList(Section)) i64 {

    // Go through each section
    var src: i64 = seed;
    for (sections.items) |sect| {

        // Find the destination corresponding to the current source
        var dst: i64 = -1;
        for (sect.ranges.items) |r| {
            if (src >= r.b and src <= r.b + r.c) {
                dst = r.a + src - r.b;
                break;
            }
        }

        // If not found, use the source value
        if (dst == -1) {
            dst = src;
        }

        // Source of next transformation is the destination of this one
        src = dst;
    }

    // Return the final destination
    return src;
}

// Read input file into list of seed numbers, and list of sections
fn read_input(fname: []const u8, allocator: std.mem.Allocator) !struct { seeds: std.ArrayList(i64), sections: std.ArrayList(Section) } {

    // Create lists for seeds and sections (caller must deinit() these)
    var seeds = std.ArrayList(i64).init(allocator);
    var sections = std.ArrayList(Section).init(allocator);

    // Read file (make sure max_bytes is big enough), any data must be copied since
    // buffer is freed at end of function
    const data = try std.fs.cwd().readFileAlloc(allocator, fname, 10000);
    defer allocator.free(data);

    // Process each line
    var it = std.mem.split(u8, data, "\n");
    while (it.next()) |l| {

        // Parse list of seed numbers
        // e.g.,: seeds: 79 14 55 13
        if (l.len > 6 and std.mem.eql(u8, l[0..6], "seeds:")) {
            try parse_nums(&seeds, l[7..]);
        }

        // Start a new section
        // e.g., seed-to-soil map:
        // Note that 'name' points to slice in data, which will be freed, so need to make a copy
        else if (std.mem.endsWith(u8, l, " map:")) {
            const name = try std.mem.Allocator.dupe(allocator, u8, l[0..(l.len - 5)]);
            const sec = Section{
                .name = name, // use copy, not slice into original string
                .ranges = std.ArrayList(Range).init(allocator),
            };
            try sections.append(sec);
        }

        // Parse a line of numbers, add it to the current section
        // e.g., 52 50 48
        else if (l.len > 0) { // Ignore blank lines
            var nums = std.ArrayList(i64).init(allocator);
            try parse_nums(&nums, l);
            if (nums.items.len == 3) {
                const r = Range{ .a = nums.items[0], .b = nums.items[1], .c = nums.items[2] };
                try sections.items[sections.items.len - 1].ranges.append(r);
            }
            nums.deinit();
        }
    }

    // Return both the seed list and the sections list, in a single anonymous structure
    return .{
        .seeds = seeds,
        .sections = sections,
    };
}

// Parse a space-separate list of numbers, putting values into passed array
fn parse_nums(a: *std.ArrayList(i64), l: []const u8) !void {
    var it = std.mem.split(u8, l, " ");
    while (it.next()) |s| {
        const n = try std.fmt.parseInt(i64, s, 10);
        try a.append(n);
    }
}
