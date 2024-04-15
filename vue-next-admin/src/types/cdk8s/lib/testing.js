"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Testing = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const fs = require("fs");
const os = require("os");
const path = require("path");
const app_1 = require("./app");
const chart_1 = require("./chart");
/**
 * Testing utilities for cdk8s applications.
 */
class Testing {
    /**
     * Returns an app for testing with the following properties:
     * - Output directory is a temp dir.
     */
    static app(props) {
        let outdir;
        if (props) {
            outdir = props.outdir ?? fs.mkdtempSync(path.join(os.tmpdir(), 'cdk8s.outdir.'));
        }
        else {
            outdir = fs.mkdtempSync(path.join(os.tmpdir(), 'cdk8s.outdir.'));
        }
        return new app_1.App({ outdir, ...props });
    }
    /**
     * @returns a Chart that can be used for tests
     */
    static chart() {
        return new chart_1.Chart(this.app(), 'test');
    }
    /**
     * Returns the Kubernetes manifest synthesized from this chart.
     */
    static synth(chart) {
        return chart.toJson();
    }
    /* istanbul ignore next */
    constructor() {
        return;
    }
}
exports.Testing = Testing;
_a = JSII_RTTI_SYMBOL_1;
Testing[_a] = { fqn: "cdk8s.Testing", version: "2.68.60" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidGVzdGluZy5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy90ZXN0aW5nLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEseUJBQXlCO0FBQ3pCLHlCQUF5QjtBQUN6Qiw2QkFBNkI7QUFDN0IsK0JBQXNDO0FBQ3RDLG1DQUFnQztBQUVoQzs7R0FFRztBQUNILE1BQWEsT0FBTztJQUNsQjs7O09BR0c7SUFDSSxNQUFNLENBQUMsR0FBRyxDQUFDLEtBQWdCO1FBQ2hDLElBQUksTUFBYyxDQUFDO1FBQ25CLElBQUksS0FBSyxFQUFFLENBQUM7WUFDVixNQUFNLEdBQUcsS0FBSyxDQUFDLE1BQU0sSUFBSSxFQUFFLENBQUMsV0FBVyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUMsRUFBRSxDQUFDLE1BQU0sRUFBRSxFQUFFLGVBQWUsQ0FBQyxDQUFDLENBQUM7UUFDbkYsQ0FBQzthQUFNLENBQUM7WUFDTixNQUFNLEdBQUcsRUFBRSxDQUFDLFdBQVcsQ0FBQyxJQUFJLENBQUMsSUFBSSxDQUFDLEVBQUUsQ0FBQyxNQUFNLEVBQUUsRUFBRSxlQUFlLENBQUMsQ0FBQyxDQUFDO1FBQ25FLENBQUM7UUFDRCxPQUFPLElBQUksU0FBRyxDQUFDLEVBQUUsTUFBTSxFQUFFLEdBQUcsS0FBSyxFQUFFLENBQUMsQ0FBQztJQUN2QyxDQUFDO0lBRUQ7O09BRUc7SUFDSSxNQUFNLENBQUMsS0FBSztRQUNqQixPQUFPLElBQUksYUFBSyxDQUFDLElBQUksQ0FBQyxHQUFHLEVBQUUsRUFBRSxNQUFNLENBQUMsQ0FBQztJQUN2QyxDQUFDO0lBRUQ7O09BRUc7SUFDSSxNQUFNLENBQUMsS0FBSyxDQUFDLEtBQVk7UUFDOUIsT0FBTyxLQUFLLENBQUMsTUFBTSxFQUFFLENBQUM7SUFDeEIsQ0FBQztJQUVELDBCQUEwQjtJQUMxQjtRQUNFLE9BQU87SUFDVCxDQUFDOztBQWhDSCwwQkFpQ0MiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgKiBhcyBmcyBmcm9tICdmcyc7XG5pbXBvcnQgKiBhcyBvcyBmcm9tICdvcyc7XG5pbXBvcnQgKiBhcyBwYXRoIGZyb20gJ3BhdGgnO1xuaW1wb3J0IHsgQXBwLCBBcHBQcm9wcyB9IGZyb20gJy4vYXBwJztcbmltcG9ydCB7IENoYXJ0IH0gZnJvbSAnLi9jaGFydCc7XG5cbi8qKlxuICogVGVzdGluZyB1dGlsaXRpZXMgZm9yIGNkazhzIGFwcGxpY2F0aW9ucy5cbiAqL1xuZXhwb3J0IGNsYXNzIFRlc3Rpbmcge1xuICAvKipcbiAgICogUmV0dXJucyBhbiBhcHAgZm9yIHRlc3Rpbmcgd2l0aCB0aGUgZm9sbG93aW5nIHByb3BlcnRpZXM6XG4gICAqIC0gT3V0cHV0IGRpcmVjdG9yeSBpcyBhIHRlbXAgZGlyLlxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBhcHAocHJvcHM/OiBBcHBQcm9wcykge1xuICAgIGxldCBvdXRkaXI6IHN0cmluZztcbiAgICBpZiAocHJvcHMpIHtcbiAgICAgIG91dGRpciA9IHByb3BzLm91dGRpciA/PyBmcy5ta2R0ZW1wU3luYyhwYXRoLmpvaW4ob3MudG1wZGlyKCksICdjZGs4cy5vdXRkaXIuJykpO1xuICAgIH0gZWxzZSB7XG4gICAgICBvdXRkaXIgPSBmcy5ta2R0ZW1wU3luYyhwYXRoLmpvaW4ob3MudG1wZGlyKCksICdjZGs4cy5vdXRkaXIuJykpO1xuICAgIH1cbiAgICByZXR1cm4gbmV3IEFwcCh7IG91dGRpciwgLi4ucHJvcHMgfSk7XG4gIH1cblxuICAvKipcbiAgICogQHJldHVybnMgYSBDaGFydCB0aGF0IGNhbiBiZSB1c2VkIGZvciB0ZXN0c1xuICAgKi9cbiAgcHVibGljIHN0YXRpYyBjaGFydCgpIHtcbiAgICByZXR1cm4gbmV3IENoYXJ0KHRoaXMuYXBwKCksICd0ZXN0Jyk7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0aGUgS3ViZXJuZXRlcyBtYW5pZmVzdCBzeW50aGVzaXplZCBmcm9tIHRoaXMgY2hhcnQuXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIHN5bnRoKGNoYXJ0OiBDaGFydCk6IGFueVtdIHtcbiAgICByZXR1cm4gY2hhcnQudG9Kc29uKCk7XG4gIH1cblxuICAvKiBpc3RhbmJ1bCBpZ25vcmUgbmV4dCAqL1xuICBwcml2YXRlIGNvbnN0cnVjdG9yKCkge1xuICAgIHJldHVybjtcbiAgfVxufVxuIl19