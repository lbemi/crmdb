"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.SizeRoundingBehavior = exports.Size = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
/**
 * Represents the amount of digital storage.
 *
 * The amount can be specified either as a literal value (e.g: `10`) which
 * cannot be negative.
 *
 * When the amount is passed as a token, unit conversion is not possible.
 */
class Size {
    /**
     * Create a Storage representing an amount kibibytes.
     * 1 KiB = 1024 bytes
     */
    static kibibytes(amount) {
        return new Size(amount, StorageUnit.Kibibytes);
    }
    /**
     * Create a Storage representing an amount mebibytes.
     * 1 MiB = 1024 KiB
     */
    static mebibytes(amount) {
        return new Size(amount, StorageUnit.Mebibytes);
    }
    /**
     * Create a Storage representing an amount gibibytes.
     * 1 GiB = 1024 MiB
     */
    static gibibytes(amount) {
        return new Size(amount, StorageUnit.Gibibytes);
    }
    /**
     * Create a Storage representing an amount tebibytes.
     * 1 TiB = 1024 GiB
     */
    static tebibytes(amount) {
        return new Size(amount, StorageUnit.Tebibytes);
    }
    /**
     * Create a Storage representing an amount pebibytes.
     * 1 PiB = 1024 TiB
     */
    static pebibyte(amount) {
        return new Size(amount, StorageUnit.Pebibytes);
    }
    constructor(amount, unit) {
        if (amount < 0) {
            throw new Error(`Storage amounts cannot be negative. Received: ${amount}`);
        }
        this.amount = amount;
        this.unit = unit;
    }
    /**
     * Return this storage as a total number of kibibytes.
     */
    toKibibytes(opts = {}) {
        return convert(this.amount, this.unit, StorageUnit.Kibibytes, opts);
    }
    /**
     * Return this storage as a total number of mebibytes.
     */
    toMebibytes(opts = {}) {
        return convert(this.amount, this.unit, StorageUnit.Mebibytes, opts);
    }
    /**
     * Return this storage as a total number of gibibytes.
     */
    toGibibytes(opts = {}) {
        return convert(this.amount, this.unit, StorageUnit.Gibibytes, opts);
    }
    /**
     * Return this storage as a total number of tebibytes.
     */
    toTebibytes(opts = {}) {
        return convert(this.amount, this.unit, StorageUnit.Tebibytes, opts);
    }
    /**
     * Return this storage as a total number of pebibytes.
     */
    toPebibytes(opts = {}) {
        return convert(this.amount, this.unit, StorageUnit.Pebibytes, opts);
    }
}
exports.Size = Size;
_a = JSII_RTTI_SYMBOL_1;
Size[_a] = { fqn: "cdk8s.Size", version: "2.68.60" };
/**
 * Rounding behaviour when converting between units of `Size`.
 */
var SizeRoundingBehavior;
(function (SizeRoundingBehavior) {
    /** Fail the conversion if the result is not an integer. */
    SizeRoundingBehavior[SizeRoundingBehavior["FAIL"] = 0] = "FAIL";
    /** If the result is not an integer, round it to the closest integer less than the result */
    SizeRoundingBehavior[SizeRoundingBehavior["FLOOR"] = 1] = "FLOOR";
    /** Don't round. Return even if the result is a fraction. */
    SizeRoundingBehavior[SizeRoundingBehavior["NONE"] = 2] = "NONE";
})(SizeRoundingBehavior || (exports.SizeRoundingBehavior = SizeRoundingBehavior = {}));
class StorageUnit {
    constructor(label, inKibiBytes) {
        this.label = label;
        this.inKibiBytes = inKibiBytes;
        // MAX_SAFE_INTEGER is 2^53, so by representing storage in kibibytes,
        // the highest storage we can represent is 8 exbibytes.
    }
    toString() {
        return this.label;
    }
}
StorageUnit.Kibibytes = new StorageUnit('kibibytes', 1);
StorageUnit.Mebibytes = new StorageUnit('mebibytes', 1024);
StorageUnit.Gibibytes = new StorageUnit('gibibytes', 1024 * 1024);
StorageUnit.Tebibytes = new StorageUnit('tebibytes', 1024 * 1024 * 1024);
StorageUnit.Pebibytes = new StorageUnit('pebibytes', 1024 * 1024 * 1024 * 1024);
function convert(amount, fromUnit, toUnit, options = {}) {
    const rounding = options.rounding ?? SizeRoundingBehavior.FAIL;
    if (fromUnit.inKibiBytes === toUnit.inKibiBytes) {
        return amount;
    }
    const multiplier = fromUnit.inKibiBytes / toUnit.inKibiBytes;
    const value = amount * multiplier;
    switch (rounding) {
        case SizeRoundingBehavior.NONE:
            return value;
        case SizeRoundingBehavior.FLOOR:
            return Math.floor(value);
        default:
        case SizeRoundingBehavior.FAIL:
            if (!Number.isInteger(value)) {
                throw new Error(`'${amount} ${fromUnit}' cannot be converted into a whole number of ${toUnit}.`);
            }
            return value;
    }
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoic2l6ZS5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9zaXplLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUE7Ozs7Ozs7R0FPRztBQUNILE1BQWEsSUFBSTtJQUNmOzs7T0FHRztJQUNJLE1BQU0sQ0FBQyxTQUFTLENBQUMsTUFBYztRQUNwQyxPQUFPLElBQUksSUFBSSxDQUFDLE1BQU0sRUFBRSxXQUFXLENBQUMsU0FBUyxDQUFDLENBQUM7SUFDakQsQ0FBQztJQUVEOzs7T0FHRztJQUNJLE1BQU0sQ0FBQyxTQUFTLENBQUMsTUFBYztRQUNwQyxPQUFPLElBQUksSUFBSSxDQUFDLE1BQU0sRUFBRSxXQUFXLENBQUMsU0FBUyxDQUFDLENBQUM7SUFDakQsQ0FBQztJQUVEOzs7T0FHRztJQUNJLE1BQU0sQ0FBQyxTQUFTLENBQUMsTUFBYztRQUNwQyxPQUFPLElBQUksSUFBSSxDQUFDLE1BQU0sRUFBRSxXQUFXLENBQUMsU0FBUyxDQUFDLENBQUM7SUFDakQsQ0FBQztJQUVEOzs7T0FHRztJQUNJLE1BQU0sQ0FBQyxTQUFTLENBQUMsTUFBYztRQUNwQyxPQUFPLElBQUksSUFBSSxDQUFDLE1BQU0sRUFBRSxXQUFXLENBQUMsU0FBUyxDQUFDLENBQUM7SUFDakQsQ0FBQztJQUVEOzs7T0FHRztJQUNJLE1BQU0sQ0FBQyxRQUFRLENBQUMsTUFBYztRQUNuQyxPQUFPLElBQUksSUFBSSxDQUFDLE1BQU0sRUFBRSxXQUFXLENBQUMsU0FBUyxDQUFDLENBQUM7SUFDakQsQ0FBQztJQUtELFlBQW9CLE1BQWMsRUFBRSxJQUFpQjtRQUNuRCxJQUFJLE1BQU0sR0FBRyxDQUFDLEVBQUUsQ0FBQztZQUNmLE1BQU0sSUFBSSxLQUFLLENBQUMsaURBQWlELE1BQU0sRUFBRSxDQUFDLENBQUM7UUFDN0UsQ0FBQztRQUNELElBQUksQ0FBQyxNQUFNLEdBQUcsTUFBTSxDQUFDO1FBQ3JCLElBQUksQ0FBQyxJQUFJLEdBQUcsSUFBSSxDQUFD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QUFLLG9CQUFvQixDQUFDLElBQUk7WUFDNUIsT0FBTyxLQUFLLENBQUM7UUFDZixLQUFLLG9CQUFvQixDQUFDLEtBQUs7WUFDN0IsT0FBTyxJQUFJLENBQUMsS0FBSyxDQUFDLEtBQUssQ0FBQyxDQUFDO1FBQzNCLFFBQVE7UUFDUixLQUFLLG9CQUFvQixDQUFDLElBQUk7WUFDNUIsSUFBSSxDQUFDLE1BQU0sQ0FBQyxTQUFTLENBQUMsS0FBSyxDQUFDLEVBQUUsQ0FBQztnQkFDN0IsTUFBTSxJQUFJLEtBQUssQ0FBQyxJQUFJLE1BQU0sSUFBSSxRQUFRLGdEQUFnRCxNQUFNLEdBQUcsQ0FBQyxDQUFDO1lBQ25HLENBQUM7WUFDRCxPQUFPLEtBQUssQ0FBQztJQUNqQixDQUFDO0FBQ0gsQ0FBQyIsInNvdXJjZXNDb250ZW50IjpbIi8qKlxuICogUmVwcmVzZW50cyB0aGUgYW1vdW50IG9mIGRpZ2l0YWwgc3RvcmFnZS5cbiAqXG4gKiBUaGUgYW1vdW50IGNhbiBiZSBzcGVjaWZpZWQgZWl0aGVyIGFzIGEgbGl0ZXJhbCB2YWx1ZSAoZS5nOiBgMTBgKSB3aGljaFxuICogY2Fubm90IGJlIG5lZ2F0aXZlLlxuICpcbiAqIFdoZW4gdGhlIGFtb3VudCBpcyBwYXNzZWQgYXMgYSB0b2tlbiwgdW5pdCBjb252ZXJzaW9uIGlzIG5vdCBwb3NzaWJsZS5cbiAqL1xuZXhwb3J0IGNsYXNzIFNpemUge1xuICAvKipcbiAgICogQ3JlYXRlIGEgU3RvcmFnZSByZXByZXNlbnRpbmcgYW4gYW1vdW50IGtpYmlieXRlcy5cbiAgICogMSBLaUIgPSAxMDI0IGJ5dGVzXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIGtpYmlieXRlcyhhbW91bnQ6IG51bWJlcik6IFNpemUge1xuICAgIHJldHVybiBuZXcgU2l6ZShhbW91bnQsIFN0b3JhZ2VVbml0LktpYmlieXRlcyk7XG4gIH1cblxuICAvKipcbiAgICogQ3JlYXRlIGEgU3RvcmFnZSByZXByZXNlbnRpbmcgYW4gYW1vdW50IG1lYmlieXRlcy5cbiAgICogMSBNaUIgPSAxMDI0IEtpQlxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBtZWJpYnl0ZXMoYW1vdW50OiBudW1iZXIpOiBTaXplIHtcbiAgICByZXR1cm4gbmV3IFNpemUoYW1vdW50LCBTdG9yYWdlVW5pdC5NZWJpYnl0ZXMpO1xuICB9XG5cbiAgLyoqXG4gICAqIENyZWF0ZSBhIFN0b3JhZ2UgcmVwcmVzZW50aW5nIGFuIGFtb3VudCBnaWJpYnl0ZXMuXG4gICAqIDEgR2lCID0gMTAyNCBNaUJcbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgZ2liaWJ5dGVzKGFtb3VudDogbnVtYmVyKTogU2l6ZSB7XG4gICAgcmV0dXJuIG5ldyBTaXplKGFtb3VudCwgU3RvcmFnZVVuaXQuR2liaWJ5dGVzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBDcmVhdGUgYSBTdG9yYWdlIHJlcHJlc2VudGluZyBhbiBhbW91bnQgdGViaWJ5dGVzLlxuICAgKiAxIFRpQiA9IDEwMjQgR2lCXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIHRlYmlieXRlcyhhbW91bnQ6IG51bWJlcik6IFNpemUge1xuICAgIHJldHVybiBuZXcgU2l6ZShhbW91bnQsIFN0b3JhZ2VVbml0LlRlYmlieXRlcyk7XG4gIH1cblxuICAvKipcbiAgICogQ3JlYXRlIGEgU3RvcmFnZSByZXByZXNlbnRpbmcgYW4gYW1vdW50IHBlYmlieXRlcy5cbiAgICogMSBQaUIgPSAxMDI0IFRpQlxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBwZWJpYnl0ZShhbW91bnQ6IG51bWJlcik6IFNpemUge1xuICAgIHJldHVybiBuZXcgU2l6ZShhbW91bnQsIFN0b3JhZ2VVbml0LlBlYmlieXRlcyk7XG4gIH1cblxuICBwcml2YXRlIHJlYWRvbmx5IGFtb3VudDogbnVtYmVyO1xuICBwcml2YXRlIHJlYWRvbmx5IHVuaXQ6IFN0b3JhZ2VVbml0O1xuXG4gIHByaXZhdGUgY29uc3RydWN0b3IoYW1vdW50OiBudW1iZXIsIHVuaXQ6IFN0b3JhZ2VVbml0KSB7XG4gICAgaWYgKGFtb3VudCA8IDApIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgU3RvcmFnZSBhbW91bnRzIGNhbm5vdCBiZSBuZWdhdGl2ZS4gUmVjZWl2ZWQ6ICR7YW1vdW50fWApO1xuICAgIH1cbiAgICB0aGlzLmFtb3VudCA9IGFtb3VudDtcbiAgICB0aGlzLnVuaXQgPSB1bml0O1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGlzIHN0b3JhZ2UgYXMgYSB0b3RhbCBudW1iZXIgb2Yga2liaWJ5dGVzLlxuICAgKi9cbiAgcHVibGljIHRvS2liaWJ5dGVzKG9wdHM6IFNpemVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBTdG9yYWdlVW5pdC5LaWJpYnl0ZXMsIG9wdHMpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGlzIHN0b3JhZ2UgYXMgYSB0b3RhbCBudW1iZXIgb2YgbWViaWJ5dGVzLlxuICAgKi9cbiAgcHVibGljIHRvTWViaWJ5dGVzKG9wdHM6IFNpemVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBTdG9yYWdlVW5pdC5NZWJpYnl0ZXMsIG9wdHMpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGlzIHN0b3JhZ2UgYXMgYSB0b3RhbCBudW1iZXIgb2YgZ2liaWJ5dGVzLlxuICAgKi9cbiAgcHVibGljIHRvR2liaWJ5dGVzKG9wdHM6IFNpemVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBTdG9yYWdlVW5pdC5HaWJpYnl0ZXMsIG9wdHMpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGlzIHN0b3JhZ2UgYXMgYSB0b3RhbCBudW1iZXIgb2YgdGViaWJ5dGVzLlxuICAgKi9cbiAgcHVibGljIHRvVGViaWJ5dGVzKG9wdHM6IFNpemVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBTdG9yYWdlVW5pdC5UZWJpYnl0ZXMsIG9wdHMpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybiB0aGlzIHN0b3JhZ2UgYXMgYSB0b3RhbCBudW1iZXIgb2YgcGViaWJ5dGVzLlxuICAgKi9cbiAgcHVibGljIHRvUGViaWJ5dGVzKG9wdHM6IFNpemVDb252ZXJzaW9uT3B0aW9ucyA9IHt9KTogbnVtYmVyIHtcbiAgICByZXR1cm4gY29udmVydCh0aGlzLmFtb3VudCwgdGhpcy51bml0LCBTdG9yYWdlVW5pdC5QZWJpYnl0ZXMsIG9wdHMpO1xuICB9XG59XG5cbi8qKlxuICogUm91bmRpbmcgYmVoYXZpb3VyIHdoZW4gY29udmVydGluZyBiZXR3ZWVuIHVuaXRzIG9mIGBTaXplYC5cbiAqL1xuZXhwb3J0IGVudW0gU2l6ZVJvdW5kaW5nQmVoYXZpb3Ige1xuICAvKiogRmFpbCB0aGUgY29udmVyc2lvbiBpZiB0aGUgcmVzdWx0IGlzIG5vdCBhbiBpbnRlZ2VyLiAqL1xuICBGQUlMLFxuICAvKiogSWYgdGhlIHJlc3VsdCBpcyBub3QgYW4gaW50ZWdlciwgcm91bmQgaXQgdG8gdGhlIGNsb3Nlc3QgaW50ZWdlciBsZXNzIHRoYW4gdGhlIHJlc3VsdCAqL1xuICBGTE9PUixcbiAgLyoqIERvbid0IHJvdW5kLiBSZXR1cm4gZXZlbiBpZiB0aGUgcmVzdWx0IGlzIGEgZnJhY3Rpb24uICovXG4gIE5PTkUsXG5cbn1cblxuLyoqXG4gKiBPcHRpb25zIGZvciBob3cgdG8gY29udmVydCB0aW1lIHRvIGEgZGlmZmVyZW50IHVuaXQuXG4gKi9cbmV4cG9ydCBpbnRlcmZhY2UgU2l6ZUNvbnZlcnNpb25PcHRpb25zIHtcbiAgLyoqXG4gICAqIEhvdyBjb252ZXJzaW9ucyBzaG91bGQgYmVoYXZlIHdoZW4gaXQgZW5jb3VudGVycyBhIG5vbi1pbnRlZ2VyIHJlc3VsdFxuICAgKiBAZGVmYXVsdCBTaXplUm91bmRpbmdCZWhhdmlvci5GQUlMXG4gICAqL1xuICByZWFkb25seSByb3VuZGluZz86IFNpemVSb3VuZGluZ0JlaGF2aW9yO1xufVxuXG5jbGFzcyBTdG9yYWdlVW5pdCB7XG4gIHB1YmxpYyBzdGF0aWMgcmVhZG9ubHkgS2liaWJ5dGVzID0gbmV3IFN0b3JhZ2VVbml0KCdraWJpYnl0ZXMnLCAxKTtcbiAgcHVibGljIHN0YXRpYyByZWFkb25seSBNZWJpYnl0ZXMgPSBuZXcgU3RvcmFnZVVuaXQoJ21lYmlieXRlcycsIDEwMjQpO1xuICBwdWJsaWMgc3RhdGljIHJlYWRvbmx5IEdpYmlieXRlcyA9IG5ldyBTdG9yYWdlVW5pdCgnZ2liaWJ5dGVzJywgMTAyNCAqIDEwMjQpO1xuICBwdWJsaWMgc3RhdGljIHJlYWRvbmx5IFRlYmlieXRlcyA9IG5ldyBTdG9yYWdlVW5pdCgndGViaWJ5dGVzJywgMTAyNCAqIDEwMjQgKiAxMDI0KTtcbiAgcHVibGljIHN0YXRpYyByZWFkb25seSBQZWJpYnl0ZXMgPSBuZXcgU3RvcmFnZVVuaXQoJ3BlYmlieXRlcycsIDEwMjQgKiAxMDI0ICogMTAyNCAqIDEwMjQpO1xuXG4gIHByaXZhdGUgY29uc3RydWN0b3IocHVibGljIHJlYWRvbmx5IGxhYmVsOiBzdHJpbmcsIHB1YmxpYyByZWFkb25seSBpbktpYmlCeXRlczogbnVtYmVyKSB7XG4gICAgLy8gTUFYX1NBRkVfSU5URUdFUiBpcyAyXjUzLCBzbyBieSByZXByZXNlbnRpbmcgc3RvcmFnZSBpbiBraWJpYnl0ZXMsXG4gICAgLy8gdGhlIGhpZ2hlc3Qgc3RvcmFnZSB3ZSBjYW4gcmVwcmVzZW50IGlzIDggZXhiaWJ5dGVzLlxuICB9XG5cbiAgcHVibGljIHRvU3RyaW5nKCkge1xuICAgIHJldHVybiB0aGlzLmxhYmVsO1xuICB9XG59XG5cbmZ1bmN0aW9uIGNvbnZlcnQoYW1vdW50OiBudW1iZXIsIGZyb21Vbml0OiBTdG9yYWdlVW5pdCwgdG9Vbml0OiBTdG9yYWdlVW5pdCwgb3B0aW9uczogU2l6ZUNvbnZlcnNpb25PcHRpb25zID0ge30pIHtcbiAgY29uc3Qgcm91bmRpbmcgPSBvcHRpb25zLnJvdW5kaW5nID8/IFNpemVSb3VuZGluZ0JlaGF2aW9yLkZBSUw7XG4gIGlmIChmcm9tVW5pdC5pbktpYmlCeXRlcyA9PT0gdG9Vbml0LmluS2liaUJ5dGVzKSB7IHJldHVybiBhbW91bnQ7IH1cblxuICBjb25zdCBtdWx0aXBsaWVyID0gZnJvbVVuaXQuaW5LaWJpQnl0ZXMgLyB0b1VuaXQuaW5LaWJpQnl0ZXM7XG4gIGNvbnN0IHZhbHVlID0gYW1vdW50ICogbXVsdGlwbGllcjtcbiAgc3dpdGNoIChyb3VuZGluZykge1xuICAgIGNhc2UgU2l6ZVJvdW5kaW5nQmVoYXZpb3IuTk9ORTpcbiAgICAgIHJldHVybiB2YWx1ZTtcbiAgICBjYXNlIFNpemVSb3VuZGluZ0JlaGF2aW9yLkZMT09SOlxuICAgICAgcmV0dXJuIE1hdGguZmxvb3IodmFsdWUpO1xuICAgIGRlZmF1bHQ6XG4gICAgY2FzZSBTaXplUm91bmRpbmdCZWhhdmlvci5GQUlMOlxuICAgICAgaWYgKCFOdW1iZXIuaXNJbnRlZ2VyKHZhbHVlKSkge1xuICAgICAgICB0aHJvdyBuZXcgRXJyb3IoYCcke2Ftb3VudH0gJHtmcm9tVW5pdH0nIGNhbm5vdCBiZSBjb252ZXJ0ZWQgaW50byBhIHdob2xlIG51bWJlciBvZiAke3RvVW5pdH0uYCk7XG4gICAgICB9XG4gICAgICByZXR1cm4gdmFsdWU7XG4gIH1cbn0iXX0=