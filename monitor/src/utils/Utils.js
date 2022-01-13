export function BigNumFormat(value, index) {
    const M = 1000000
    const B = M * 1000
    const T = B * 1000
    if (M <= value && value < B) {
        value = value / M + " M";
    } else if (B <= value && value < T) {
        value = value / B + " B";
    } else if (T <= value) {
        value = value / T + " T";
    }
    return value
}
