import {render, screen} from "@testing-library/svelte";
import {expect, test} from "vitest";
import Page from "./+page.svelte";
import {mockGetSummaryResponse} from "$lib/utils/tests/summaryFactory";

const data = {
    summary: mockGetSummaryResponse()
};


test("/daily/+page.svelte", async () => {
    // @ts-ignore
    render(Page, {data});

    const displayDate = screen.getByTestId("dailyDate");
    expect(displayDate).toBeInTheDocument();

    const winnerName = screen.getByTestId("dailyWinnerName");
    expect(winnerName).toBeInTheDocument();

    const previous = screen.getByTestId("dailyPreviousButton");
    expect(previous).toBeInTheDocument();

    const previousIcon = screen.getByTestId("dailyPreviousIcon");
    expect(previousIcon).toBeInTheDocument();
    expect(previous).toContainElement(previousIcon);

    const next = screen.getByTestId("dailyNextButton");
    expect(next).toBeInTheDocument();

    const goalCarousel = screen.getByTestId("dailyGoalCarousel");
    expect(goalCarousel).toBeInTheDocument();

    const goalItem = screen.getByTestId("dailyGoal-0");
    expect(goalItem).toBeInTheDocument();
    expect(goalCarousel).toContainElement(goalItem);

    const workoutList = screen.getByTestId("dailyWorkoutsList");
    expect(workoutList).toBeInTheDocument();

    const workoutItem = screen.getByTestId("dailyWorkoutItem-0");
    expect(workoutItem).toBeInTheDocument();
    expect(workoutList).toContainElement(workoutItem);

});
