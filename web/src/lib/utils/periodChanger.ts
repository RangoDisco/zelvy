export const handlePeriodChange = (rawSD: string, rawED: string, direction: "previous" | "next"): {
    startDate: string,
    endDate: string
} => {
    let endDate = new Date(rawED);

    endDate.setMonth(direction === "next" ? endDate.getMonth() + 1 : endDate.getMonth() - 1);

    if (new Date() < endDate) {
        endDate = new Date();
    }

    const startDate = new Date(endDate.getFullYear(), endDate.getMonth() - 1);

    return {
        endDate: endDate.toISOString(),
        startDate: startDate.toISOString()
    };
};
