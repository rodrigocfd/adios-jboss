'use strict';

const {exec} = require('child_process');

function runc(cmd) {
	return new Promise((resolve, reject) => {
		exec(cmd, (err, stdout, stderr) => {
			resolve(stdout.split('\n').filter(line => line.length > 0));
		});
	});
}

(async () => {
	const psGrep = await runc('ps aux | grep jboss');
	let jbossFound = false;

	for (const line of psGrep) {
		if (line.includes('/bin/java')) {
			jbossFound = true;
			const tokens = line.split(' ').filter(s => s !== '');
			const killStr = 'kill -9 ' + tokens[1];
			console.log('Tocando o foda-se: ' + killStr);
			await runc(killStr);
		}
	}

	if (!jbossFound) {
		console.log('Esta bosta não está rodando.');
	}
})();
