import type {WinnerViewModel} from "$lib/gen/zelvy/user/winner_view_model";
import type {HeatmapItemViewModel} from "$lib/gen/zelvy/summary/heatmap_item_view_model";

export const formatWinner = (winners: WinnerViewModel[]): OverviewStat => {
    const winner = winners[0];
    return {
        picto: "👑",
        title: winner.username,
        subtitle: "Most wins",
        value: winner.wins
    };
};

export const formatSuccessRate = (days: HeatmapItemViewModel[]): OverviewStat => {
    const totalDays = days.length;
    const successNumber = days.filter((day: HeatmapItemViewModel) => day.success).length;
    const successRate = (successNumber / totalDays * 100).toFixed(2);

    return {
        picto: "🎯",
        title: `${successRate}%`,
        subtitle: "Success Rate",
        value: `${successNumber}/${totalDays}`
    };
};

export const formatLongestStreak = (days: HeatmapItemViewModel[]): OverviewStat => {
    let longestStreak = 0;

    days.forEach(day => {
        if (!day.success) {
            longestStreak = 0;
        } else {
            longestStreak++;
        }
    });

    return {
        picto: "⛓️‍💥",
        title: `${longestStreak} days`,
        subtitle: "Longest Streak",
        value: `Out of ${days.length}`
    };
};
