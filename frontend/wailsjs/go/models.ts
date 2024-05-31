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
	
	export class Issue {
	    self: string;
	    id: number;
	
	    static createFrom(source: any = {}) {
	        return new Issue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.self = source["self"];
	        this.id = source["id"];
	    }
	}
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
		    if (a.slice && a.map) {
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
	export class WorkLogAttr {
	    self: string;
	    key: string;
	    id: number;
	    name: string;
	    status: string;
	    global: boolean;
	
	    static createFrom(source: any = {}) {
	        return new WorkLogAttr(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.self = source["self"];
	        this.key = source["key"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.status = source["status"];
	        this.global = source["global"];
	    }
	}
	export class WorkLogValue {
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new WorkLogValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class WorkLogAttribute {
	    self: string;
	    values: WorkLogValue[];
	
	    static createFrom(source: any = {}) {
	        return new WorkLogAttribute(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.self = source["self"];
	        this.values = this.convertValues(source["values"], WorkLogValue);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class WorkLogResult {
	    self: string;
	    tempoWorklogId: number;
	    issue: Issue;
	    timeSpentSeconds: number;
	    billableSeconds: number;
	    startDate: string;
	    startTime: string;
	    description: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    attributes: WorkLogAttribute;
	
	    static createFrom(source: any = {}) {
	        return new WorkLogResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.self = source["self"];
	        this.tempoWorklogId = source["tempoWorklogId"];
	        this.issue = this.convertValues(source["issue"], Issue);
	        this.timeSpentSeconds = source["timeSpentSeconds"];
	        this.billableSeconds = source["billableSeconds"];
	        this.startDate = source["startDate"];
	        this.startTime = source["startTime"];
	        this.description = source["description"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.attributes = this.convertValues(source["attributes"], WorkLogAttribute);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

