#!/usr/bin/env node
/** 
 * obelisk.js is an HTML Xpath parser to get 6502 opcode data from http://www.obelisk.me.uk/6502/reference.html 
 */

/** @todo 
0x1A, 0x3A, 0x5A, 0x7A, 0xDA, 0xFA: // implied undocumented NOPs (1 byte wide)
0x0C, 0x1C, 0x3C, 0x5C, 0x7C, 0xDC, 0xFC: cpu.PC += 2 // absolute undoc NOPs (3 bytes wide)
0x80, 0x82, 0x89, 0xC2, 0xE2, 0x04, 0x14, 0x34, 0x44, 0x54, 0x64, 0x74, 0xD4, 0xF4: cpu.PC++ // immediate and zeropage undoc NOPs (2 bytes wide)
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

console.log( "export default "+ JSON.stringify(opcodes, null, "\t") +";");
