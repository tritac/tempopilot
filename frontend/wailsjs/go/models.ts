export namespace appstore {
	
	export class UserConfig {
	    api_key: string;
	    user_name: string;
	    user_id: string;
	    isValidApi: boolean;
	
	    static createFrom(source: any = {}) {
	        return new UserConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.api_key = source["api_key"];
	        this.user_name = source["user_name"];
	        this.user_id = source["user_id"];
	        this.isValidApi = source["isValidApi"];
	    }
	}

}

export namespace worklog {
	
	export class WorkDay {
	    // Go type: time
	    date: any;
	    day: string;
	    isWorking: boolean;
	
	    static createFrom(source: any = {}) {
	        return new WorkDay(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = this.convertValues(source["date"], null);
	        this.day = source["day"];
	        this.isWorking = source["isWorking"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

