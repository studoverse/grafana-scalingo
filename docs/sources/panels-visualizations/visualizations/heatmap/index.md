---
aliases:
  - ../../features/panels/heatmap/
  - ../../visualizations/heatmap/
description: Configure options for Grafana's heatmap visualization
keywords:
  - grafana
  - heatmap
  - panel
  - documentation
labels:
  products:
    - cloud
    - enterprise
    - oss
title: Heatmap
weight: 100
---

# Heatmap

Heatmaps allow you to view [histograms](https://grafana.com/docs/grafana/latest/panels-visualizations/visualizations/histogram/) over time. While histograms display the data distribution that falls in a specific value range, heatmaps allow you to identify patterns in the histogram data distribution over time. For more information about heatmaps, refer to [Introduction to histograms and heatmaps](https://grafana.com/docs/grafana/latest/fundamentals/intro-histograms/).

For example, if you want to understand the temperature changes for the past few years, you can use a heatmap visualization to identify trends in your data:

{{< figure src="/static/img/docs/heatmap-panel/temperature_heatmap.png" max-width="1025px" alt="A heatmap visualization showing the random walk distribution over time" >}}

You can use a heatmap visualization if you need to:

- Visualize a large density of your data distribution.
- Condense large amounts of data through various color schemes that are easier to interpret.
- Identify any outliers in your data distribution.
- Provide statistical analysis to see how values or trends change over time.

## Configure a heatmap visualization

Once you’ve created a [dashboard](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/dashboards/build-dashboards/create-dashboard/), the following video shows you how to configure a heatmap visualization:

{{< youtube id="SGWBzQ54koE" >}}

## Supported data formats

Heatmaps support time series data.

### Example

The table below is a simplified output of random walk distribution over time:

| Time                | Walking (km) |
| ------------------- | ------------ |
| 2023-06-25 21:13:09 | 10           |
| 2023-08-25 21:13:10 | 8            |
| 2023-08-30 21:13:10 | 10           |
| 2023-10-08 21:13:11 | 12           |
| 2023-12-25 21:13:11 | 14           |
| 2024-01-05 21:13:12 | 13           |
| 2024-02-22 21:13:13 | 10           |

The data is converted as follows:

{{< figure src="/static/img/docs/heatmap-panel/heatmap.png" max-width="1025px" alt="A heatmap visualization showing the random walk distribution over time" >}}

## Heatmap options

### Calculate from data

This setting determines if the data is already a calculated heatmap (from the data source/transformer), or one that should be calculated in the panel.

### X Bucket

This setting determines how the X-axis is split into buckets. You can specify a time interval in the **Size** input. For example, a time range of `1h` makes the cells 1-hour wide on the X-axis.

### Y Bucket

This setting determines how the Y-axis is split into buckets.

### Y Bucket scale

Select one of the following Y-axis value scales:

- **linear -** Linear scale.
- **log (base 2) -** Logarithmic scale with base 2.
- **log (base 10) -** Logarithmic scale with base 10.
- **symlog -** Symlog scale.

## Y Axes

Defines how the Y axis is displayed

### Placement

- **Left** On the left
- **Right** On the right
- **Hidden** Hidden

### Unit

Unit configuration

### Decimals

This setting determines decimal configuration.

### Min/Max value

This setting configures the axis range.

### Axis width

This setting configures the width for the axis.

### Axis value

This setting configures the axis value.

### Reverse

When selected, the axis appears in reverse order.

{{< docs/shared lookup="visualizations/multiple-y-axes.md" source="grafana" version="<GRAFANA VERSION>" leveloffset="+2" >}}

## Colors

The color spectrum controls the mapping between value count (in each bucket) and the color assigned to each bucket. The leftmost color on the spectrum represents the minimum count and the color on the right most side represents the maximum count. Some color schemes are automatically inverted when using the light theme.

You can also change the color mode to Opacity. In this case, the color will not change but the amount of opacity will change with the bucket count

- **Mode**
  - **Scheme -** Bucket value represented by cell color.
    - **Scheme -** If the mode is **scheme**, then select a color scheme.
  - **opacity -** Bucket value represented by cell opacity. Opaque cell means maximum value.
    - **Color -** Cell base color.
    - **Scale -** Scale for mapping bucket values to the opacity.
      - **linear -** Linear scale. Bucket value maps linearly to the opacity.
      - **sqrt -** Power scale. Cell opacity calculated as `value ^ k`, where `k` is a configured **Exponent** value. If exponent is less than `1`, you will get a logarithmic scale. If exponent is greater than `1`, you will get an exponential scale. In case of `1`, scale will be the same as linear.
    - **Exponent -** value of the exponent, greater than `0`.

### Start/end color from value

By default, Grafana calculates cell colors based on minimum and maximum bucket values. With Min and Max you can overwrite those values. Consider a bucket value as a Z-axis and Min and Max as Z-Min and Z-Max, respectively.

- **Start -** Minimum value using for cell color calculation. If the bucket value is less than Min, then it is mapped to the "minimum" color. The series min value is the default value.
- **End -** Maximum value using for cell color calculation. If the bucket value is greater than Max, then it is mapped to the "maximum" color. The series max value is the default value.

## Cell display

Use these settings to refine your visualization.

## Additional display options

### Tooltip

- **Show tooltip -** Show heatmap tooltip.
- **Show Histogram -** Show a Y-axis histogram on the tooltip. A histogram represents the distribution of the bucket values for a specific timestamp.
- **Show color scale -** Show a color scale on the tooltip. The color scale represents the mapping between bucket value and color. This option is configurable when you enable the `newVizTooltips` feature flag.

### Legend

Choose whether you want to display the heatmap legend on the visualization.

### Exemplars

Set the color used to show exemplar data.

{{% docs/reference %}}
[Introduction to histograms and heatmaps]: "/docs/grafana/ -> /docs/grafana/<GRAFANA VERSION>/fundamentals/intro-histograms"
[Introduction to histograms and heatmaps]: "/docs/grafana-cloud/ -> /docs/grafana/<GRAFANA VERSION>/fundamentals/intro-histograms"
{{% /docs/reference %}}
