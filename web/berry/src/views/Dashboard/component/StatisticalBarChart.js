import PropTypes from 'prop-types';

// material-ui
import { Grid, Typography, FormControl, Select, MenuItem } from '@mui/material';

// third-party
import Chart from 'react-apexcharts';

// project imports
import SkeletonTotalGrowthBarChart from 'ui-component/cards/Skeleton/TotalGrowthBarChart';
import MainCard from 'ui-component/cards/MainCard';
import { gridSpacing } from 'store/constant';
import { Box } from '@mui/material';

// ==============================|| DASHBOARD DEFAULT - TOTAL GROWTH BAR CHART ||============================== //

const StatisticalBarChart = ({ isLoading, chartDatas, statisticsMetric, onMetricChange }) => {
  // Update chart data with current metric-specific formatting
  const updatedChartData = {
    ...chartData,
    options: {
      ...chartData.options,
      xaxis: {
        ...chartData.options.xaxis,
        categories: chartDatas.xaxis
      },
      tooltip: {
        ...chartData.options.tooltip,
        y: {
          formatter: function (val) {
            switch (statisticsMetric) {
              case 'requests':
                return val.toLocaleString();
              case 'tokens':
                return val.toLocaleString();
              case 'expenses':
              default:
                return '$' + val;
            }
          }
        }
      }
    },
    series: chartDatas.data
  };

  return (
    <>
      {isLoading ? (
        <SkeletonTotalGrowthBarChart />
      ) : (
        <MainCard>
          <Grid container spacing={gridSpacing}>
            <Grid item xs={12}>
              <Grid container alignItems="center" justifyContent="space-between">
                <Grid item>
                  <Typography variant="h3">
                    统计 - {
                      statisticsMetric === 'tokens' ? '令牌' :
                      statisticsMetric === 'requests' ? '请求' :
                      '费用'
                    }
                  </Typography>
                </Grid>
                <Grid item>
                  <FormControl size="small" sx={{ minWidth: 120 }}>
                    <Select
                      value={statisticsMetric}
                      onChange={(e) => onMetricChange(e, e.target.value)}
                      displayEmpty
                    >
                      <MenuItem value="expenses">费用</MenuItem>
                      <MenuItem value="requests">请求</MenuItem>
                      <MenuItem value="tokens">令牌</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
              </Grid>
            </Grid>
            <Grid item xs={12}>
              {updatedChartData.series ? (
                <Chart {...updatedChartData} />
              ) : (
                <Box
                  sx={{
                    minHeight: '490px',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center'
                  }}
                >
                  <Typography variant="h3" color={'#697586'}>
                    暂无数据
                  </Typography>
                </Box>
              )}
            </Grid>
          </Grid>
        </MainCard>
      )}
    </>
  );
};

StatisticalBarChart.propTypes = {
  isLoading: PropTypes.bool,
  chartDatas: PropTypes.object,
  statisticsMetric: PropTypes.string,
  onMetricChange: PropTypes.func
};

export default StatisticalBarChart;

const chartData = {
  height: 480,
  type: 'bar',
  options: {
    colors: [
      '#008FFB',
      '#00E396',
      '#FEB019',
      '#FF4560',
      '#775DD0',
      '#55efc4',
      '#81ecec',
      '#74b9ff',
      '#a29bfe',
      '#00b894',
      '#00cec9',
      '#0984e3',
      '#6c5ce7',
      '#ffeaa7',
      '#fab1a0',
      '#ff7675',
      '#fd79a8',
      '#fdcb6e',
      '#e17055',
      '#d63031',
      '#e84393'
    ],
    chart: {
      id: 'bar-chart',
      stacked: true,
      toolbar: {
        show: true
      },
      zoom: {
        enabled: true
      }
    },
    responsive: [
      {
        breakpoint: 480,
        options: {
          legend: {
            position: 'bottom',
            offsetX: -10,
            offsetY: 0
          }
        }
      }
    ],
    plotOptions: {
      bar: {
        horizontal: false,
        columnWidth: '50%'
      }
    },
    xaxis: {
      type: 'category',
      categories: []
    },
    legend: {
      show: true,
      fontSize: '14px',
      fontFamily: `'Roboto', sans-serif`,
      position: 'bottom',
      offsetX: 20,
      labels: {
        useSeriesColors: false
      },
      markers: {
        width: 16,
        height: 16,
        radius: 5
      },
      itemMargin: {
        horizontal: 15,
        vertical: 8
      }
    },
    fill: {
      type: 'solid'
    },
    dataLabels: {
      enabled: false
    },
    grid: {
      show: true
    },
    tooltip: {
      theme: 'dark',
      fixed: {
        enabled: false
      },
      y: {
        formatter: function (val) {
          return '$' + val;
        }
      },
      marker: {
        show: false
      }
    }
  },
  series: []
};
