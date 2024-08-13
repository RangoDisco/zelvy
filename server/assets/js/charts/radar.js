const createRadar = (radar, ctx) => {
  return new Chart(ctx, {
    type: "radar",
    data: {
      labels: radar.labels,
      datasets: buildRadarSet(radar.datasets),
    },
    options: {
      layout: {
        padding: 10,
      },
      scales: {
        r: {
          ticks: {
            display: false,
          },
          angleLines: {
            display: false,
            color: "rgba(255, 255, 255, 0.2)",
          },
          grid: {
            color: "rgba(255, 255, 255, 0.2)",
          },
          pointLabels: {
            font: 12,
            color: "rgba(255, 255, 255, 0.5)",
          },
        },
      },
      plugins: {
        legend: {
          position: "bottom",
          labels: {
            color: "rgba(255, 255, 255, 0.5)",
          },
        },
      },
    },
  });
};

const buildRadarSet = (rawDatasets) => {
  const sets = [];
  rawDatasets.forEach((set) => {
    sets.push({
      label: set.label,
      data: set.data,
      fill: true,
      backgroundColor: set.backgroundColor,
      borderColor: set.borderColor,
      borderWidth: 1,
      pointBackgroundColor: set.pointBackgroundColor,
      pointBorderColor: "#fff",
    });
  });

  return sets;
};
