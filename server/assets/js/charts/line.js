const createLine = (line, ctx) => {
  return new Chart(ctx, {
    type: "line",
    data: {
      labels: line.labels,
      datasets: buildLineSet(line.datasets),
    },
    options: {
      layout: {
        padding: 10,
      },
      scales: {
        x: {
          grid: {
            color: "rgba(255, 255, 255, 0.3)",
          },

          ticks: {
            color: "rgba(255, 255, 255, 0.5)",
          },
        },
        y: {
          grid: {
            color: "rgba(255, 255, 255, 0.3)",
          },

          ticks: {
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

const buildLineSet = (rawDatasets) => {
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
