///////////////////////////////////////////////////////////
//  synchronous callback based push-fashion reactivity   //
///////////////////////////////////////////////////////////

class Dependency {
	constructor(state) {
		this.state = state;
		this.update = null;
		this.subscribers = new Set();
	}
	dependent() {
		if (this.update) {
			this.subscribers.add(this.update);
		}
		this.update = null;
	}
	notify() {
		this.subscribers.forEach((update) => {
			update();
		});
	}
	onStateChange(update) {
		this.update = update;
		this.dependent();
		// initial udpate
		update();
	}
	set(cb) {
		this.state = cb(this.state);
		this.notify();
	}
}

let pow2 = 0;
let pow3 = 0;
let dep = new Dependency(3);
dep.onStateChange(() => {
	pow2 = dep.state * dep.state;
});
dep.onStateChange(() => {
	pow3 = dep.state * dep.state * dep.state;
});
console.log(`dep ${dep.state}: pow2 ${pow2} pow3 ${pow3}`);

dep.set((state) => {
	return state + 1;
});
console.log(`dep ${dep.state}: pow2 ${pow2} pow3 ${pow3}`);
