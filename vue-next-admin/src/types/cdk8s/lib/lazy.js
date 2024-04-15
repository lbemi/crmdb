"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Lazy = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
class Lazy {
    static any(producer) {
        return new Lazy(producer);
    }
    constructor(producer) {
        this.producer = producer;
    }
    produce() {
        return this.producer.produce();
    }
}
exports.Lazy = Lazy;
_a = JSII_RTTI_SYMBOL_1;
Lazy[_a] = { fqn: "cdk8s.Lazy", version: "2.68.60" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibGF6eS5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9sYXp5LnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEsTUFBYSxJQUFJO0lBQ1IsTUFBTSxDQUFDLEdBQUcsQ0FBQyxRQUFzQjtRQUN0QyxPQUFPLElBQUksSUFBSSxDQUFDLFFBQVEsQ0FBbUIsQ0FBQztJQUM5QyxDQUFDO0lBRUQsWUFBcUMsUUFBc0I7UUFBdEIsYUFBUSxHQUFSLFFBQVEsQ0FBYztJQUFHLENBQUM7SUFFeEQsT0FBTztRQUNaLE9BQU8sSUFBSSxDQUFDLFFBQVEsQ0FBQyxPQUFPLEVBQUUsQ0FBQztJQUNqQyxDQUFDOztBQVRILG9CQVVDIiwic291cmNlc0NvbnRlbnQiOlsiZXhwb3J0IGNsYXNzIExhenkge1xuICBwdWJsaWMgc3RhdGljIGFueShwcm9kdWNlcjogSUFueVByb2R1Y2VyKTogYW55IHtcbiAgICByZXR1cm4gbmV3IExhenkocHJvZHVjZXIpIGFzIHVua25vd24gYXMgYW55O1xuICB9XG5cbiAgcHJpdmF0ZSBjb25zdHJ1Y3Rvcihwcml2YXRlIHJlYWRvbmx5IHByb2R1Y2VyOiBJQW55UHJvZHVjZXIpIHt9XG5cbiAgcHVibGljIHByb2R1Y2UoKTogYW55IHtcbiAgICByZXR1cm4gdGhpcy5wcm9kdWNlci5wcm9kdWNlKCk7XG4gIH1cbn1cblxuZXhwb3J0IGludGVyZmFjZSBJQW55UHJvZHVjZXIge1xuICBwcm9kdWNlKCk6IGFueTtcbn1cblxuXG4iXX0=