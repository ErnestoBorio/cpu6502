#!/usr/bin/env node
/** 
 * obelisk.js is an HTML Xpath parser to get 6502 opcode data from http://www.obelisk.me.uk/6502/reference.html 
 */

import { select as x } from "xpath";
import { DOMParser } from "xmldom";
import fs from "fs";

const xml = fs.readFileSync(process.argv[2] || "./6502 Reference.html", "utf8");
const dom = new DOMParser({
	errorHandler: {
		warning: () => null,
		error: (e) => null,
		fatalError: (e) => console.error(e),
	},
}).parseFromString(xml);

let opcodes = Array(0x100);

const ops = x("//h3[a[@name]]", dom); /** <h3><a name="LDA">LDA - Load Accumulator</a></h3> */

let debug;

try {
	ops.forEach((op, i) => {
		const mnemonic = x("a/@name", op)[0].nodeValue;
		const table = x("following-sibling::p[position() <= 5]/table/tbody", op)[1];
		const rows = x("tr[position() > 1]", table);
		rows.forEach((row) => {
			debug = {mnemonic, row, x: x("td[3]/center/text()", row)};
			const addressing = x("td[1]/a/@href", row)[0].nodeValue.split("#")[1];
			const addressingLong = x("td[1]/a/text()", row)[0].nodeValue.trim();
			const opcode = x("td[2]/center/text()", row)[0].nodeValue.trim().slice(1,3);
			const opcodeNum = parseInt("0x"+opcode, 16);
			const bytes = parseInt( x("td[3]/center/text()", row)[0].nodeValue.trim(), 10);
			let cycleText = "";
			x("td[4]/text()", row).forEach((node) => cycleText += node.nodeValue);
			const cycles = parseInt(cycleText.slice(0,1), 10);
			opcodes[opcodeNum] = {
				opcodeNum,
				opcode,
				mnemonic,
				addressing,
				addressingLong,
				bytes,
				cycles,
			};
			const extra = cycleText.length > 1 ? cycleText.slice(2) : undefined;
			if(extra) {
				opcodes[parseInt("0x"+opcode, 16)].extra = extra;
			}
		});
	});
}
catch(err) {
	console.error(err);
}

console.log( JSON.stringify(opcodes, null, "\t"));