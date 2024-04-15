"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Duration = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
/**
 * Represents a length of time.
 *
 * The amount can be specified either as a literal value (e.g: `10`) which
 * cannot be negative.
 *
 */
class Duration {
    /**
     * Create a Duration representing an amount of milliseconds
     *
     * @param amount the amount of Milliseconds the `Duration` will represent.
     * @returns a new `Duration` representing `amount` ms.
     */
    static millis(amount) {
        return new Duration(amount, TimeUnit.Milliseconds);
    }
    /**
     * Create a Duration representing an amount of seconds
     *
     * @param amount the amount of Seconds the `Duration` will represent.
     * @returns a new `Duration` representing `amount` Seconds.
     */
    static seconds(amount) {
        return new Duration(amount, TimeUnit.Seconds);
    }
    /**
     * Create a Duration representing an amount of minutes
     *
     * @param amount the amount of Minutes the `Duration` will represent.
     * @returns a new `Duration` representing `amount` Minutes.
     */
    static minutes(amount) {
        return new Duration(amount, TimeUnit.Minutes);
    }
    /**
     * Create a Duration representing an amount of hours
     *
     * @param amount the amount of Hours the `Duration` will represent.
     * @returns a new `Duration` representing `amount` Hours.
     */
    static hours(amount) {
        return new Duration(amount, TimeUnit.Hours);
    }
    /**
     * Create a Duration representing an amount of days
     *
     * @param amount the amount of Days the `Duration` will represent.
     * @returns a new `Duration` representing `amount` Days.
     */
    static days(amount) {
        return new Duration(amount, TimeUnit.Days);
    }
    /**
     * Parse a period formatted according to the ISO 8601 standard
     *
     * @see https://www.iso.org/fr/standard/70907.html
     * @param duration an ISO-formtted duration to be parsed.
     * @returns the parsed `Duration`.
     */
    static parse(duration) {
        const matches = duration.match(/^PT(?:(\d+)D)?(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?$/);
        if (!matches) {
            throw new Error(`Not a valid ISO duration: ${duration}`);
        }
        const [, days, hours, minutes, seconds] = matches;
        if (!days && !hours && !minutes && !seconds) {
            throw new Error(`Not a valid ISO duration: ${duration}`);
        }
        return Duration.millis(_toInt(seconds) * TimeUnit.Seconds.inMillis
            + (_toInt(minutes) * TimeUnit.Minutes.inMillis)
            + (_toInt(hours) * TimeUnit.Hours.inMillis)
            + (_toInt(days) * TimeUnit.Days.inMillis));
        function _toInt(str) {
            if (!str) {
                return 0;
            }
            return Number(str);
        }
    }
    constructor(amount, unit) {
        if (amount < 0) {
            throw new Error(`Duration amounts cannot be negative. Received: ${amount}`);
        }
        this.amount = amount;
        this.unit = unit;
    }
    /**
     * Return the total number of milliseconds in this Duration
     *
     * @returns the value of this `Duration` expressed in Milliseconds.
     */
    toMilliseconds(opts = {}) {
        return convert(this.amount, this.unit, TimeUnit.Milliseconds, opts);
    }
    /**
     * Return the total number of seconds in this Duration
     *
     * @returns the value of this `Duration` expressed in Seconds.
     */
    toSeconds(opts = {}) {
        return convert(this.amount, this.unit, TimeUnit.Seconds, opts);
    }
    /**
     * Return the total number of minutes in this Duration
     *
     * @returns the value of this `Duration` expressed in Minutes.
     */
    toMinutes(opts = {}) {
        return convert(this.amount, this.unit, TimeUnit.Minutes, opts);
    }
    /**
     * Return the total number of hours in this Duration
     *
     * @returns the value of this `Duration` expressed in Hours.
     */
    toHours(opts = {}) {
        return convert(this.amount, this.unit, TimeUnit.Hours, opts);
    }
    /**
     * Return the total number of days in this Duration
     *
     * @returns the value of this `Duration` expressed in Days.
     */
    toDays(opts = {}) {
        return convert(this.amount, this.unit, TimeUnit.Days, opts);
    }
    /**
     * Return an ISO 8601 representation of this period
     *
     * @returns a string starting with 'PT' describing the period
     * @see https://www.iso.org/fr/standard/70907.html
     */
    toIsoString() {
        if (this.amount === 0) {
            return 'PT0S';
        }
        switch (this.unit) {
            case TimeUnit.Seconds: return `PT${this.fractionDuration('S', 60, Duration.minutes)}`;
            case TimeUnit.Minutes: return `PT${this.fractionDuration('M', 60, Duration.hours)}`;
            case TimeUnit.Hours: return `PT${this.fractionDuration('H', 24, Duration.days)}`;
            case TimeUnit.Days: return `PT${this.amount}D`;
            default:
                throw new Error(`Unexpected time unit: ${this.unit}`);
        }
    }
    /**
     * Turn this duration into a human-readable string
     */
    toHumanString() {
        if (this.amount === 0) {
            return fmtUnit(0, this.unit);
        }
        let millis = convert(this.amount, this.unit, TimeUnit.Milliseconds, { integral: false });
        const parts = new Array();
        for (const unit of [TimeUnit.Days, TimeUnit.Hours, TimeUnit.Hours, TimeUnit.Minutes, TimeUnit.Seconds]) {
            const wholeCount = Math.floor(convert(millis, TimeUnit.Milliseconds, unit, { integral: false }));
            if (wholeCount > 0) {
                parts.push(fmtUnit(wholeCount, unit));
                millis -= wholeCount * unit.inMillis;
            }
        }
        // Remainder in millis
        if (millis > 0) {
            parts.push(fmtUnit(millis, TimeUnit.Milliseconds));
        }
        // 2 significant parts, that's totally enough for humans
        return parts.slice(0, 2).join(' ');
        function fmtUnit(amount, unit) {
            if (amount === 1) {
                // All of the labels end in 's'
                return `${amount} ${unit.label.substring(0, unit.label.length - 1)}`;
            }
            return `${amount} ${unit.label}`;
        }
    }
    /**
     * Return unit of Duration
     */
    unitLabel() {
        return this.unit.toString();
    }
    fractionDuration(symbol, modulus, next) {
        if (this.amount < modulus) {
            return `${this.amount}${symbol}`;
        }
        const remainder = this.amount % modulus;
        const quotient = next((this.amount - remainder) / modulus).toIsoString().slice(2);
        return remainder > 0
            ? `${quotient}${remainder}${symbol}`
            : quotient;
    }
}
exports.Duration = Duration;
_a = JSII_RTTI_SYMBOL_1;
Duration[_a] = { fqn: "cdk8s.Duration", version: "2.68.60" };
class TimeUnit {
    constructor(label, inMillis) {
        this.label = label;
        this.inMillis = inMillis;
        // MAX_SAFE_INTEGER is 2^53, so by representing our duration in millis (the lowest
        // common unit) the highest duration we can represent is
        // 2^53 / 86*10^6 ~= 104 * 10^6 days (about 100 million days).
    }
    toString() {
        return this.label;
    }
}
TimeUnit.Milliseconds = new TimeUnit('millis', 1);
TimeUnit.Seconds = new TimeUnit('seconds', 1000);
TimeUnit.Minutes = new TimeUnit('minutes', 60000);
TimeUnit.Hours = new TimeUnit('hours', 3600000);
TimeUnit.Days = new TimeUnit('days', 86400000);
function convert(amount, fromUnit, toUnit, { integral = true }) {
    if (fromUnit.inMillis === toUnit.inMillis) {
        if (integral && !Number.isInteger(amount)) {
            throw new Error(`${amount} must be a whole number of ${toUnit}.`);
        }
        return amount;
    }
    const multiplier = fromUnit.inMillis / toUnit.inMillis;
    const value = amount * multiplier;
    if (!Number.isInteger(value) && integral) {
        throw new Error(`'${amount} ${fromUnit}' cannot be converted into a whole number of ${toUnit}.`);
    }
    return value;
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZHVyYXRpb24uanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvZHVyYXRpb24udHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7QUFBQTs7Ozs7O0dBTUc7QUFDSCxNQUFhLFFBQVE7SUFDbkI7Ozs7O09BS0c7SUFDSSxNQUFNLENBQUMsTUFBTSxDQUFDLE1BQWM7UUFDakMsT0FBTyxJQUFJLFFBQVEsQ0FBQyxNQUFNLEVBQUUsUUFBUSxDQUFDLFlBQVksQ0FBQyxDQUFDO0lBQ3JELENBQUM7SUFFRDs7Ozs7T0FLRztJQUNJLE1BQU0sQ0FBQyxPQUFPLENBQUMsTUFBYztRQUNsQyxPQUFPLElBQUksUUFBUSxDQUFDLE1BQU0sRUFBRSxRQUFRLENBQUMsT0FBTyxDQUFDLENBQUM7SUFDaEQsQ0FBQztJQUVEOzs7OztPQUtHO0lBQ0ksTUFBTSxDQUFDLE9BQU8sQ0FBQyxNQUFjO1FBQ2xDLE9BQU8sSUFBSSxRQUFRLENBQUMsTUFBTSxFQUFFLFFBQVEsQ0FBQyxPQUFPLENBQUMsQ0FBQztJQUNoRCxDQUFDO0lBRUQ7Ozs7O09BS0c7SUFDSSxNQUFNLENBQUMsS0FBSyxDQUFDLE1BQWM7UUFDaEMsT0FBTyxJQUFJLFFBQVEsQ0FBQyxNQUFNLEVBQUUsUUFBUSxDQUFDLEtBQUssQ0FBQyxDQUFDO0lBQzlDLENBQUM7SUFFRDs7Ozs7T0FLRztJQUNJLE1BQU0sQ0FBQyxJQUFJLENBQUMsTUFBYztRQUMvQixPQUFPLElBQUksUUFBUSxDQUFDLE1BQU0sRUFBRSxRQUFRLENBQUMsSUFBSSxDQUFDLENBQUM7SUFDN0MsQ0FBQztJQUVEOzs7Ozs7T0FNRztJQUNJLE1BQU0sQ0FBQyxLQUFLLENBQUMsUUFBZ0I7UUFDbEMsTUFBTSxPQUFPLEdBQUcsUUFBUSxDQUFDLEtBQUssQ0FBQyxrREFBa0QsQ0FBQyxDQUFDO1FBQ25GLElBQUksQ0FBQyxPQUFPLEVBQUUsQ0FBQztZQUNiLE1BQU0sSUFBSSxLQUFLLENBQUMsNkJBQTZCLFFBQVEsRUFBRSxD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QUFNLEVBQUUsUUFBVSxDQUFDLENBQUM7QUFhakUsU0FBUyxPQUFPLENBQUMsTUFBYyxFQUFFLFFBQWtCLEVBQUUsTUFBZ0IsRUFBRSxFQUFFLFFBQVEsR0FBRyxJQUFJLEVBQXlCO0lBQy9HLElBQUksUUFBUSxDQUFDLFFBQVEsS0FBSyxNQUFNLENBQUMsUUFBUSxFQUFFLENBQUM7UUFDMUMsSUFBSSxRQUFRLElBQUksQ0FBQyxNQUFNLENBQUMsU0FBUyxDQUFDLE1BQU0sQ0FBQyxFQUFFLENBQUM7WUFDMUMsTUFBTSxJQUFJLEtBQUssQ0FBQyxHQUFHLE1BQU0sOEJBQThCLE1BQU0sR0FBRyxDQUFDLENBQUM7UUFDcEUsQ0FBQztRQUNELE9BQU8sTUFBTSxDQUFDO0lBQ2hCLENBQUM7SUFDRCxNQUFNLFVBQVUsR0FBRyxRQUFRLENBQUMsUUFBUSxHQUFHLE1BQU0sQ0FBQyxRQUFRLENBQUM7SUFFdkQsTUFBTSxLQUFLLEdBQUcsTUFBTSxHQUFHLFVBQVUsQ0FBQztJQUNsQyxJQUFJLENBQUMsTUFBTSxDQUFDLFNBQVMsQ0FBQyxLQUFLLENBQUMsSUFBSSxRQUFRLEVBQUUsQ0FBQztRQUN6QyxNQUFNLElBQUksS0FBSyxDQUFDLElBQUksTUFBTSxJQUFJLFFBQVEsZ0RBQWdELE1BQU0sR0FBRyxDQUFDLENBQUM7SUFDbkcsQ0FBQztJQUNELE9BQU8sS0FBSyxDQUFDO0FBQ2YsQ0FBQyIsInNvdXJjZXNDb250ZW50IjpbIi8qKlxuICogUmVwcmVzZW50cyBhIGxlbmd0aCBvZiB0aW1lLlxuICpcbiAqIFRoZSBhbW91bnQgY2FuIGJlIHNwZWNpZmllZCBlaXRoZXIgYXMgYSBsaXRlcmFsIHZhbHVlIChlLmc6IGAxMGApIHdoaWNoXG4gKiBjYW5ub3QgYmUgbmVnYXRpdmUuXG4gKlxuICovXG5leHBvcnQgY2xhc3MgRHVyYXRpb24ge1xuICAvKipcbiAgICogQ3JlYXRlIGEgRHVyYXRpb24gcmVwcmVzZW50aW5nIGFuIGFtb3VudCBvZiBtaWxsaXNlY29uZHNcbiAgICpcbiAgICogQHBhcmFtIGFtb3VudCB0aGUgYW1vdW50IG9mIE1pbGxpc2Vjb25kcyB0aGUgYER1cmF0aW9uYCB3aWxsIHJlcHJlc2VudC5cbiAgICogQHJldHVybnMgYSBuZXcgYER1cmF0aW9uYCByZXByZXNlbnRpbmcgYGFtb3VudGAgbXMuXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIG1pbGxpcyhhbW91bnQ6IG51bWJlcik6IER1cmF0aW9uIHtcbiAgICByZXR1cm4gbmV3IER1cmF0aW9uKGFtb3VudCwgVGltZVVuaXQuTWlsbGlzZWNvbmRzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBDcmVhdGUgYSBEdXJhdGlvbiByZXByZXNlbnRpbmcgYW4gYW1vdW50IG9mIHNlY29uZHNcbiAgICpcbiAgICogQHBhcmFtIGFtb3VudCB0aGUgYW1vdW50IG9mIFNlY29uZHMgdGhlIGBEdXJhdGlvbmAgd2lsbCByZXByZXNlbnQuXG4gICAqIEByZXR1cm5zIGEgbmV3IGBEdXJhdGlvbmAgcmVwcmVzZW50aW5nIGBhbW91bnRgIFNlY29uZHMuXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIHNlY29uZHMoYW1vdW50OiBudW1iZXIpOiBEdXJhdGlvbiB7XG4gICAgcmV0dXJuIG5ldyBEdXJhdGlvbihhbW91bnQsIFRpbWVVbml0LlNlY29uZHMpO1xuICB9XG5cbiAgLyoqXG4gICAqIENyZWF0ZSBhIER1cmF0aW9uIHJlcHJlc2VudGluZyBhbiBhbW91bnQgb2YgbWludXRlc1xuICAgKlxuICAgKiBAcGFyYW0gYW1vdW50IHRoZSBhbW91bnQgb2YgTWludXRlcyB0aGUgYER1cmF0aW9uYCB3aWxsIHJlcHJlc2VudC5cbiAgICogQHJldHVybnMgYSBuZXcgYER1cmF0aW9uYCByZXByZXNlbnRpbmcgYGFtb3VudGAgTWludXRlcy5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgbWludXRlcyhhbW91bnQ6IG51bWJlcik6IER1cmF0aW9uIHtcbiAgICByZXR1cm4gbmV3IER1cmF0aW9uKGFtb3VudCwgVGltZVVuaXQuTWludXRlcyk7XG4gIH1cblxuICAvKipcbiAgICogQ3JlYXRlIGEgRHVyYXRpb24gcmVwcmVzZW50aW5nIGFuIGFtb3VudCBvZiBob3Vyc1xuICAgKlxuICAgKiBAcGFyYW0gYW1vdW50IHRoZSBhbW91bnQgb2YgSG91cnMgdGhlIGBEdXJhdGlvbmAgd2lsbCByZXByZXNlbnQuXG4gICAqIEByZXR1cm5zIGEgbmV3IGBEdXJhdGlvbmAgcmVwcmVzZW50aW5nIGBhbW91bnRgIEhvdXJzLlxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBob3VycyhhbW91bnQ6IG51bWJlcik6IER1cmF0aW9uIHtcbiAgICByZXR1cm4gbmV3IER1cmF0aW9uKGFtb3VudCwgVGltZVVuaXQuSG91cnMpO1xuICB9XG5cbiAgLyoqXG4gICAqIENyZWF0ZSBhIER1cmF0aW9uIHJlcHJlc2VudGluZyBhbiBhbW91bnQgb2YgZGF5c1xuICAgKlxuICAgKiBAcGFyYW0gYW1vdW50IHRoZSBhbW91bnQgb2YgRGF5cyB0aGUgYER1cmF0aW9uYCB3aWxsIHJlcHJlc2VudC5cbiAgICogQHJldHVybnMgYSBuZXcgYER1cmF0aW9uYCByZXByZXNlbnRpbmcgYGFtb3VudGAgRGF5cy5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgZGF5cyhhbW91bnQ6IG51bWJlcik6IER1cmF0aW9uIHtcbiAgICByZXR1cm4gbmV3IER1cmF0aW9uKGFtb3VudCwgVGltZVVuaXQuRGF5cyk7XG4gIH1cblxuICAvKipcbiAgICogUGFyc2UgYSBwZXJpb2QgZm9ybWF0dGVkIGFjY29yZGluZyB0byB0aGUgSVNPIDg2MDEgc3RhbmRhcmRcbiAgICpcbiAgICogQHNlZSBodHRwczovL3d3dy5pc28ub3JnL2ZyL3N0YW5kYXJkLzcwOTA3Lmh0bWxcbiAgICogQHBhcmFtIGR1cmF0aW9uIGFuIElTTy1mb3JtdHRlZCBkdXJhdGlvbiB0byBiZSBwYXJzZWQuXG4gICAqIEByZXR1cm5zIHRoZSBwYXJzZWQgYER1cmF0aW9uYC5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgcGFyc2UoZHVyYXRpb246IHN0cmluZyk6IER1cmF0aW9uIHtcbiAgICBjb25zdCBtYXRjaGVzID0gZHVyYXRpb24ubWF0Y2goL15QVCg/OihcXGQrKUQpPyg/OihcXGQrKUgpPyg/OihcXGQrKU0pPyg/OihcXGQrKVMpPyQvKTtcbiAgICBpZiAoIW1hdGNoZXMpIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgTm90IGEgdmFsaWQgSVNPIGR1cmF0aW9uOiAke2R1cmF0aW9ufWApO1xuICAgIH1cbiAgICBjb25zdCBbLCBkYXlzLCBob3VycywgbWludXRlcywgc2Vjb25kc10gPSBtYXRjaGVzO1xuICAgIGlmICghZGF5cyAmJiAhaG91cnMgJiYgIW1pbnV0ZXMgJiYgIXNlY29uZHMpIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgTm90IGEgdmFsaWQgSVNPIGR1cmF0aW9uOiAke2R1cmF0aW9ufWApO1xuICAgIH1cbiAgICByZXR1cm4gRHVyYXRpb24ubWlsbGlzKFxuICAgICAgX3RvSW50KHNlY29uZHMpICogVGltZVVuaXQuU2Vjb25kcy5pbk1pbGxpc1xuICAgICAgKyAoX3RvSW50KG1pbnV0ZXMpICogVGltZVVuaXQuTWludXRlcy5pbk1pbGxpcylcbiAgICAgICsgKF90b0ludChob3VycykgKiBUaW1lVW5pdC5Ib3Vycy5pbk1pbGxpcylcbiAgICAgICsgKF90b0ludChkYXlzKSAqIFRpbWVVbml0LkRheXMuaW5NaWxsaXMpLFxuICAgICk7XG5cbiAgICBmdW5jdGlvbiBfdG9JbnQoc3RyOiBzdHJpbmcpOiBudW1iZXIge1xuICAgICAgaWYgKCFzdHIpIHsgcmV0dXJuIDA7IH1cbiAgICAgIHJldHVybiBOdW1iZXIoc3RyKTtcbiAgICB9XG4gIH1cblxuICBwcml2YXRlIHJlYWRvbmx5IGFtb3VudDogbnVtYmVyO1xuICBwcml2YXRlIHJlYWRvbmx5IHVuaXQ6IFRpbWVVbml0O1xuXG4gIHByaXZhdGUgY29uc3RydWN0b3IoYW1vdW50OiBudW1iZXIsIHVuaXQ6IFRpbWVVbml0KSB7XG4gICAgaWYgKGFtb3VudCA8IDApIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgRHVyYXRpb24gYW1vdW50cyBjYW5ub3QgYmUgbmVnYXRpdmUuIFJlY2VpdmVkOiAke2Ftb3VudH1gKTtcbiAgICB9XG5cbiAgICB0aGlzLmFtb3VudCA9IGFtb3VudDtcbiAgICB0aGlzLnVuaXQgPSB1bml0O1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGUgdG90YWwgbnVtYmVyIG9mIG1pbGxpc2Vjb25kcyBpbiB0aGlzIER1cmF0aW9uXG4gICAqXG4gICAqIEByZXR1cm5zIHRoZSB2YWx1ZSBvZiB0aGlzIGBEdXJhdGlvbmAgZXhwcmVzc2VkIGluIE1pbGxpc2Vjb25kcy5cbiAgICovXG4gIHB1YmxpYyB0b01pbGxpc2Vjb25kcyhvcHRzOiBUaW1lQ29udmVyc2lvbk9wdGlvbnMgPSB7fSk6IG51bWJlciB7XG4gICAgcmV0dXJuIGNvbnZlcnQodGhpcy5hbW91bnQsIHRoaXMudW5pdCwgVGltZVVuaXQuTWlsbGlzZWNvbmRzLCBvcHRzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBSZXR1cm4gdGhlIHRvdGFsIG51bWJlciBvZiBzZWNvbmRzIGluIHRoaXMgRHVyYXRpb25cbiAgICpcbiAgICogQHJldHVybnMgdGhlIHZhbHVlIG9mIHRoaXMgYER1cmF0aW9uYCBleHByZXNzZWQgaW4gU2Vjb25kcy5cbiAgICovXG4gIHB1YmxpYyB0b1NlY29uZHMob3B0czogVGltZUNvbnZlcnNpb25PcHRpb25zID0ge30pOiBudW1iZXIge1xuICAgIHJldHVybiBjb252ZXJ0KHRoaXMuYW1vdW50LCB0aGlzLnVuaXQsIFRpbWVVbml0LlNlY29uZHMsIG9wdHMpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGUgdG90YWwgbnVtYmVyIG9mIG1pbnV0ZXMgaW4gdGhpcyBEdXJhdGlvblxuICAgKlxuICAgKiBAcmV0dXJucyB0aGUgdmFsdWUgb2YgdGhpcyBgRHVyYXRpb25gIGV4cHJlc3NlZCBpbiBNaW51dGVzLlxuICAgKi9cbiAgcHVibGljIHRvTWludXRlcyhvcHRzOiBUaW1lQ29udmVyc2lvbk9wdGlvbnMgPSB7fSk6IG51bWJlciB7XG4gICAgcmV0dXJuIGNvbnZlcnQodGhpcy5hbW91bnQsIHRoaXMudW5pdCwgVGltZVVuaXQuTWludXRlcywgb3B0cyk7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJuIHRoZSB0b3RhbCBudW1iZXIgb2YgaG91cnMgaW4gdGhpcyBEdXJhdGlvblxuICAgKlxuICAgKiBAcmV0dXJucyB0aGUgdmFsdWUgb2YgdGhpcyBgRHVyYXRpb25gIGV4cHJlc3NlZCBpbiBIb3Vycy5cbiAgICovXG4gIHB1YmxpYyB0b0hvdXJzKG9wdHM6IFRpbWVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBUaW1lVW5pdC5Ib3Vycywgb3B0cyk7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJuIHRoZSB0b3RhbCBudW1iZXIgb2YgZGF5cyBpbiB0aGlzIER1cmF0aW9uXG4gICAqXG4gICAqIEByZXR1cm5zIHRoZSB2YWx1ZSBvZiB0aGlzIGBEdXJhdGlvbmAgZXhwcmVzc2VkIGluIERheXMuXG4gICAqL1xuICBwdWJsaWMgdG9EYXlzKG9wdHM6IFRpbWVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBUaW1lVW5pdC5EYXlzLCBvcHRzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBSZXR1cm4gYW4gSVNPIDg2MDEgcmVwcmVzZW50YXRpb24gb2YgdGhpcyBwZXJpb2RcbiAgICpcbiAgICogQHJldHVybnMgYSBzdHJpbmcgc3RhcnRpbmcgd2l0aCAnUFQnIGRlc2NyaWJpbmcgdGhlIHBlcmlvZFxuICAgKiBAc2VlIGh0dHBzOi8vd3d3Lmlzby5vcmcvZnIvc3RhbmRhcmQvNzA5MDcuaHRtbFxuICAgKi9cbiAgcHVibGljIHRvSXNvU3RyaW5nKCk6IHN0cmluZyB7XG4gICAgaWYgKHRoaXMuYW1vdW50ID09PSAwKSB7IHJldHVybiAnUFQwUyc7IH1cbiAgICBzd2l0Y2ggKHRoaXMudW5pdCkge1xuICAgICAgY2FzZSBUaW1lVW5pdC5TZWNvbmRzOiByZXR1cm4gYFBUJHt0aGlzLmZyYWN0aW9uRHVyYXRpb24oJ1MnLCA2MCwgRHVyYXRpb24ubWludXRlcyl9YDtcbiAgICAgIGNhc2UgVGltZVVuaXQuTWludXRlczogcmV0dXJuIGBQVCR7dGhpcy5mcmFjdGlvbkR1cmF0aW9uKCdNJywgNjAsIER1cmF0aW9uLmhvdXJzKX1gO1xuICAgICAgY2FzZSBUaW1lVW5pdC5Ib3VyczogcmV0dXJuIGBQVCR7dGhpcy5mcmFjdGlvbkR1cmF0aW9uKCdIJywgMjQsIER1cmF0aW9uLmRheXMpfWA7XG4gICAgICBjYXNlIFRpbWVVbml0LkRheXM6IHJldHVybiBgUFQke3RoaXMuYW1vdW50fURgO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgdGhyb3cgbmV3IEVycm9yKGBVbmV4cGVjdGVkIHRpbWUgdW5pdDogJHt0aGlzLnVuaXR9YCk7XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIFR1cm4gdGhpcyBkdXJhdGlvbiBpbnRvIGEgaHVtYW4tcmVhZGFibGUgc3RyaW5nXG4gICAqL1xuICBwdWJsaWMgdG9IdW1hblN0cmluZygpOiBzdHJpbmcge1xuICAgIGlmICh0aGlzLmFtb3VudCA9PT0gMCkgeyByZXR1cm4gZm10VW5pdCgwLCB0aGlzLnVuaXQpOyB9XG5cbiAgICBsZXQgbWlsbGlzID0gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBUaW1lVW5pdC5NaWxsaXNlY29uZHMsIHsgaW50ZWdyYWw6IGZhbHNlIH0pO1xuICAgIGNvbnN0IHBhcnRzID0gbmV3IEFycmF5PHN0cmluZz4oKTtcblxuICAgIGZvciAoY29uc3QgdW5pdCBvZiBbVGltZVVuaXQuRGF5cywgVGltZVVuaXQuSG91cnMsIFRpbWVVbml0LkhvdXJzLCBUaW1lVW5pdC5NaW51dGVzLCBUaW1lVW5pdC5TZWNvbmRzXSkge1xuICAgICAgY29uc3Qgd2hvbGVDb3VudCA9IE1hdGguZmxvb3IoY29udmVydChtaWxsaXMsIFRpbWVVbml0Lk1pbGxpc2Vjb25kcywgdW5pdCwgeyBpbnRlZ3JhbDogZmFsc2UgfSkpO1xuICAgICAgaWYgKHdob2xlQ291bnQgPiAwKSB7XG4gICAgICAgIHBhcnRzLnB1c2goZm10VW5pdCh3aG9sZUNvdW50LCB1bml0KSk7XG4gICAgICAgIG1pbGxpcyAtPSB3aG9sZUNvdW50ICogdW5pdC5pbk1pbGxpcztcbiAgICAgIH1cbiAgICB9XG5cbiAgICAvLyBSZW1haW5kZXIgaW4gbWlsbGlzXG4gICAgaWYgKG1pbGxpcyA+IDApIHtcbiAgICAgIHBhcnRzLnB1c2goZm10VW5pdChtaWxsaXMsIFRpbWVVbml0Lk1pbGxpc2Vjb25kcykpO1xuICAgIH1cblxuICAgIC8vIDIgc2lnbmlmaWNhbnQgcGFydHMsIHRoYXQncyB0b3RhbGx5IGVub3VnaCBmb3IgaHVtYW5zXG4gICAgcmV0dXJuIHBhcnRzLnNsaWNlKDAsIDIpLmpvaW4oJyAnKTtcblxuICAgIGZ1bmN0aW9uIGZtdFVuaXQoYW1vdW50OiBudW1iZXIsIHVuaXQ6IFRpbWVVbml0KSB7XG4gICAgICBpZiAoYW1vdW50ID09PSAxKSB7XG4gICAgICAgIC8vIEFsbCBvZiB0aGUgbGFiZWxzIGVuZCBpbiAncydcbiAgICAgICAgcmV0dXJuIGAke2Ftb3VudH0gJHt1bml0LmxhYmVsLnN1YnN0cmluZygwLCB1bml0LmxhYmVsLmxlbmd0aCAtIDEpfWA7XG4gICAgICB9XG4gICAgICByZXR1cm4gYCR7YW1vdW50fSAke3VuaXQubGFiZWx9YDtcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJuIHVuaXQgb2YgRHVyYXRpb25cbiAgICovXG4gIHB1YmxpYyB1bml0TGFiZWwoKTogc3RyaW5nIHtcbiAgICByZXR1cm4gdGhpcy51bml0LnRvU3RyaW5nKCk7XG4gIH1cblxuICBwcml2YXRlIGZyYWN0aW9uRHVyYXRpb24oc3ltYm9sOiBzdHJpbmcsIG1vZHVsdXM6IG51bWJlciwgbmV4dDogKGFtb3VudDogbnVtYmVyKSA9PiBEdXJhdGlvbik6IHN0cmluZyB7XG4gICAgaWYgKHRoaXMuYW1vdW50IDwgbW9kdWx1cykge1xuICAgICAgcmV0dXJuIGAke3RoaXMuYW1vdW50fSR7c3ltYm9sfWA7XG4gICAgfVxuICAgIGNvbnN0IHJlbWFpbmRlciA9IHRoaXMuYW1vdW50ICUgbW9kdWx1cztcbiAgICBjb25zdCBxdW90aWVudCA9IG5leHQoKHRoaXMuYW1vdW50IC0gcmVtYWluZGVyKSAvIG1vZHVsdXMpLnRvSXNvU3RyaW5nKCkuc2xpY2UoMik7XG4gICAgcmV0dXJuIHJlbWFpbmRlciA+IDBcbiAgICAgID8gYCR7cXVvdGllbnR9JHtyZW1haW5kZXJ9JHtzeW1ib2x9YFxuICAgICAgOiBxdW90aWVudDtcbiAgfVxufVxuXG4vKipcbiAqIE9wdGlvbnMgZm9yIGhvdyB0byBjb252ZXJ0IHRpbWUgdG8gYSBkaWZmZXJlbnQgdW5pdC5cbiAqL1xuZXhwb3J0IGludGVyZmFjZSBUaW1lQ29udmVyc2lvbk9wdGlvbnMge1xuICAvKipcbiAgICogSWYgYHRydWVgLCBjb252ZXJzaW9ucyBpbnRvIGEgbGFyZ2VyIHRpbWUgdW5pdCAoZS5nLiBgU2Vjb25kc2AgdG8gYE1pbnV0ZXNgKSB3aWxsIGZhaWwgaWYgdGhlIHJlc3VsdCBpcyBub3QgYW5cbiAgICogaW50ZWdlci5cbiAgICpcbiAgICogQGRlZmF1bHQgdHJ1ZVxuICAgKi9cbiAgcmVhZG9ubHkgaW50ZWdyYWw/OiBib29sZWFuO1xufVxuXG5jbGFzcyBUaW1lVW5pdCB7XG4gIHB1YmxpYyBzdGF0aWMgcmVhZG9ubHkgTWlsbGlzZWNvbmRzID0gbmV3IFRpbWVVbml0KCdtaWxsaXMnLCAxKTtcbiAgcHVibGljIHN0YXRpYyByZWFkb25seSBTZWNvbmRzID0gbmV3IFRpbWVVbml0KCdzZWNvbmRzJywgMV8wMDApO1xuICBwdWJsaWMgc3RhdGljIHJlYWRvbmx5IE1pbnV0ZXMgPSBuZXcgVGltZVVuaXQoJ21pbnV0ZXMnLCA2MF8wMDApO1xuICBwdWJsaWMgc3RhdGljIHJlYWRvbmx5IEhvdXJzID0gbmV3IFRpbWVVbml0KCdob3VycycsIDNfNjAwXzAwMCk7XG4gIHB1YmxpYyBzdGF0aWMgcmVhZG9ubHkgRGF5cyA9IG5ldyBUaW1lVW5pdCgnZGF5cycsIDg2XzQwMF8wMDApO1xuXG4gIHByaXZhdGUgY29uc3RydWN0b3IocHVibGljIHJlYWRvbmx5IGxhYmVsOiBzdHJpbmcsIHB1YmxpYyByZWFkb25seSBpbk1pbGxpczogbnVtYmVyKSB7XG4gICAgLy8gTUFYX1NBRkVfSU5URUdFUiBpcyAyXjUzLCBzbyBieSByZXByZXNlbnRpbmcgb3VyIGR1cmF0aW9uIGluIG1pbGxpcyAodGhlIGxvd2VzdFxuICAgIC8vIGNvbW1vbiB1bml0KSB0aGUgaGlnaGVzdCBkdXJhdGlvbiB3ZSBjYW4gcmVwcmVzZW50IGlzXG4gICAgLy8gMl41MyAvIDg2KjEwXjYgfj0gMTA0ICogMTBeNiBkYXlzIChhYm91dCAxMDAgbWlsbGlvbiBkYXlzKS5cbiAgfVxuXG4gIHB1YmxpYyB0b1N0cmluZygpIHtcbiAgICByZXR1cm4gdGhpcy5sYWJlbDtcbiAgfVxufVxuXG5mdW5jdGlvbiBjb252ZXJ0KGFtb3VudDogbnVtYmVyLCBmcm9tVW5pdDogVGltZVVuaXQsIHRvVW5pdDogVGltZVVuaXQsIHsgaW50ZWdyYWwgPSB0cnVlIH06IFRpbWVDb252ZXJzaW9uT3B0aW9ucykge1xuICBpZiAoZnJvbVVuaXQuaW5NaWxsaXMgPT09IHRvVW5pdC5pbk1pbGxpcykge1xuICAgIGlmIChpbnRlZ3JhbCAmJiAhTnVtYmVyLmlzSW50ZWdlcihhbW91bnQpKSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoYCR7YW1vdW50fSBtdXN0IGJlIGEgd2hvbGUgbnVtYmVyIG9mICR7dG9Vbml0fS5gKTtcbiAgICB9XG4gICAgcmV0dXJuIGFtb3VudDtcbiAgfVxuICBjb25zdCBtdWx0aXBsaWVyID0gZnJvbVVuaXQuaW5NaWxsaXMgLyB0b1VuaXQuaW5NaWxsaXM7XG5cbiAgY29uc3QgdmFsdWUgPSBhbW91bnQgKiBtdWx0aXBsaWVyO1xuICBpZiAoIU51bWJlci5pc0ludGVnZXIodmFsdWUpICYmIGludGVncmFsKSB7XG4gICAgdGhyb3cgbmV3IEVycm9yKGAnJHthbW91bnR9ICR7ZnJvbVVuaXR9JyBjYW5ub3QgYmUgY29udmVydGVkIGludG8gYSB3aG9sZSBudW1iZXIgb2YgJHt0b1VuaXR9LmApO1xuICB9XG4gIHJldHVybiB2YWx1ZTtcbn0iXX0=