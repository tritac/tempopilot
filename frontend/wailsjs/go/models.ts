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

