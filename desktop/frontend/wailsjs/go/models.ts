export namespace store {
	
	export class Fiber {
	    ts: string;
	    data: number[];
	    mimetype: string;
	
	    static createFrom(source: any = {}) {
	        return new Fiber(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ts = source["ts"];
	        this.data = source["data"];
	        this.mimetype = source["mimetype"];
	    }
	}

}

