import {render, screen} from "@testing-library/svelte";
import {expect, test} from "vitest";
import Page from "./+page.svelte";
import {mockGetWinnerResponse} from "$lib/utils/tests/winnerFactory";
import {mockGetHeatmapResponse} from "$lib/utils/tests/summaryFactory";
import {mockGetWorkoutsResponse} from "$lib/utils/tests/workoutFactory";

const data = {
    winRes: mockGetWinnerResponse(),
    hmRes: mockGetHeatmapResponse(),
    wkrRes: mockGetWorkoutsResponse()
};


test("/overview/+page.svelte", async () => {
    // @ts-ignore
    render(Page, {data});

    const period = screen.getByTestId("overviewPeriod");
    expect(period).toBeInTheDocument();

    const previous = screen.getByTestId("overviewPreviousButton");
    expect(previous).toBeInTheDocument();

    const previousIcon = screen.getByTestId("overviewPreviousIcon");
    expect(previousIcon).toBeInTheDocument();
    expect(previous).toContainElement(previousIcon);

    const next = screen.getByTestId("overviewNextButton");
    expect(next).toBeInTheDocument();

    const winnerCard = screen.getByTestId("overviewWinnerCard");
    expect(winnerCard).toBeInTheDocument();

    const successRateCard = screen.getByTestId("overviewSuccessRateCard");
    expect(successRateCard).toBeInTheDocument();

    const longestStreakCard = screen.getByTestId("overviewLongestStreakCard");
    expect(longestStreakCard).toBeInTheDocument();

    const heatmapGrid = screen.getByTestId("overviewHeatmapGrid");
    expect(heatmapGrid).toBeInTheDocument();

    const heatmapGridItem = screen.getByTestId("overviewHeatmapGridItem-0");
    expect(heatmapGridItem).toBeInTheDocument();
    expect(heatmapGrid).toContainElement(heatmapGridItem);

    const workoutBarChart = screen.getByTestId("overviewBarChart");
    expect(workoutBarChart).toBeInTheDocument();
});
