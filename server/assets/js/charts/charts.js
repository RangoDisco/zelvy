const renderCharts = (charts, document) => {
  charts.forEach((chartData, index) => {
    const ctx = document.getElementById(`chart-${index}`);
    renderIndividualChart(chartData, ctx);
  });
};

const renderIndividualChart = (chartData, ctx) => {
  switch (chartData.type) {
    case "radar":
      createRadar(chartData, ctx);
    // case "line":
    //   createLine(chartData, ctx);
    default:
      throw new Error("Invalid chart type");
  }
};
