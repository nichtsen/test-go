///////////////////////////////////////////////////////////
//  synchronous callback based push-fashion reactivity   //
///////////////////////////////////////////////////////////

class Dependency {
	constructor(state) {
		this.update = null;
		this.subscribers = new Set();
		this.state = state;

		let notify = () => {
			this.subscribers.forEach((update) => {
				update();
			});
		};
		Object.keys(state).forEach((key) => {
			let value = state[key];
			Object.defineProperty(state, key, {
				get() {
					return value;
				},
				set(v) {
					value = v;
					notify();
				},
				enumerable: true,
				configurable: true,
			});
		});
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
let effect = (dep, extract) => {
	dep.onStateChange(() => {
		let count = extract(dep);
		pow2 = count * count;
	});
	dep.onStateChange(() => {
		let count = extract(dep);
		pow3 = count * count * count;
	});
};

let depOfObject = () => {
	let state = {
		count: 3,
	};
	let dep = new Dependency(state);
	effect(dep, (dep) => {
		return dep.state.count;
	});
	console.log(`dep ${dep.state.count}: pow2 ${pow2} pow3 ${pow3}`);

	state.count++;
	console.log(`dep ${dep.state.count}: pow2 ${pow2} pow3 ${pow3}`);
};
let depOfNumber = () => {
	pow2 = 0;
	pow3 = 0;
	let dep = new Dependency(3);
	effect(dep, (dep) => {
		return dep.state;
	});
	console.log(`dep ${dep.state}: pow2 ${pow2} pow3 ${pow3}`);

	dep.set((state) => {
		return state + 1;
	});

	console.log(`dep ${dep.state}: pow2 ${pow2} pow3 ${pow3}`);
};
depOfObject();
depOfNumber();
// Output:
// dep 3: pow2 9 pow3 27
// dep 4: pow2 16 pow3 64
// dep 3: pow2 9 pow3 27
// dep 4: pow2 16 pow3 64
