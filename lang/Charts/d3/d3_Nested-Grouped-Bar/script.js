// Code goes here
var margin = {
    top: 20,
    right: 20,
    bottom: 60,
    left: 40
  },
  width = 560 - margin.left - margin.right,
  height = 360 - margin.top - margin.bottom;

var x_groups = d3.scale.ordinal()
  .rangeRoundBands([0, width], .1);

var x_categories = d3.scale.ordinal();

var x_values = d3.scale.ordinal();

var y = d3.scale.linear()
  .range([height, 0]);

var color = d3.scale.ordinal()
  .range(["#98abc5", "#8a89a6", "#7b6888", "#6b486b", "#a05d56", "#d0743c", "#ff8c00"]);

var groups_axis = d3.svg.axis()
  .scale(x_groups)
  .orient("bottom");

var categories_axis = d3.svg.axis()
  .scale(x_categories)
  .orient("bottom");

var values_axis = d3.svg.axis()
  .scale(x_categories)
  .orient("bottom");

var yAxis = d3.svg.axis()
  .scale(y)
  .orient("left")
  .tickFormat(d3.format(".2s"));

var svg = d3.select("body").append("svg")
  .attr("width", width + margin.left + margin.right)
  .attr("height", height + margin.top + margin.bottom)
  .append("g")
  .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

d3.csv("data.csv", function(error, data) {
  var nested = d3.nest()
    .key(function(d) {
      return d.groups;
    })
    .key(function(d) {
      return d.categories;
    })
    .rollup(function(leaves) {
      return [{
        key: 'v-a',
        value: leaves[0]['value 1']
      }, {
        key: 'v-b',
        value: leaves[0]['value 2']
      }, {
        key: 'v-c',
        value: leaves[0]['value 3']
      }];
    })
    .entries(data);

  x_groups.domain(nested.map(function(d) {
    return d.key;
  }));
  //var categories = ['A', 'B', 'C']; 
  var categories = nested[0].values.map(function(d, i) {
    return d.key;
  });
  x_categories.domain(categories).rangeRoundBands([0, x_groups.rangeBand()]);
  //var values = ['value 1', 'value 2', 'value 3']; 
  var values = nested[0].values[0].values.map(function(d, i) {
    return d.key;
  });
  x_values.domain(values).rangeRoundBands([0, x_categories.rangeBand()]);

  // ew! should prob do something way cleaner
  y.domain([0, d3.max(nested, function(d) {
    return d3.max(d.values, function(d) {
      return d3.max(d.values, function(d) {
        return d.value;
      })
    });
  })])

  svg.append("g")
    .attr("class", "x axis")
    .attr("transform", "translate(0," + (height + 30) + ")")
    .call(groups_axis);

  svg.append("g")
    .attr("class", "y axis")
    .call(yAxis)
    .append("text")
    .attr("transform", "rotate(-90)")
    .attr("y", 6)
    .attr("dy", ".71em")
    .style("text-anchor", "end")
    .text("Value");

  var groups_g = svg.selectAll(".group")
    .data(nested)
    .enter().append("g")
    .attr("class", function(d) {
      return 'group group-' + d.key;
    })
    .attr("transform", function(d) {
      return "translate(" + x_groups(d.key) + ",0)";
    });

  var categories_g = groups_g.selectAll(".category")
    .data(function(d) {
      return d.values;
    })
    .enter().append("g")
    .attr("class", function(d) {
      return 'category category-' + d.key;
    })
    .attr("transform", function(d) {
      return "translate(" + x_categories(d.key) + ",0)";
    });
    
  var categories_labels = categories_g.selectAll('.category-label')
    .data(function(d) {
      return [d.key];
    })
    .enter().append("text")
    .attr("class", function(d) {
      return 'category-label category-label-' + d;
    })
    .attr("x", function(d) {
      return x_categories.rangeBand() / 2;
    })
    .attr('y', function(d) {
      return height + 25;
    })
    .attr('text-anchor', 'middle')
    .text(function(d) {
      return d;
    })
    
  var values_g = categories_g.selectAll(".value")
    .data(function(d) {
      return d.values;
    })
    .enter().append("g")
    .attr("class", function(d) {
      return 'value value-' + d.key;
    })
    .attr("transform", function(d) {
      return "translate(" + x_values(d.key) + ",0)";
    });

  var values_labels = values_g.selectAll('.value-label')
    .data(function(d) {
      return [d.key];
    })
    .enter().append("text")
    .attr("class", function(d) {
      return 'value-label value-label-' + d;
    })
    .attr("x", function(d) {
      return x_values.rangeBand() / 2;
    })
    .attr('y', function(d) {
      return height + 10;
    })
    .attr('text-anchor', 'middle')
    .text(function(d) {
      return d;
    })

  var rects = values_g.selectAll('.rect')
    .data(function(d) {
      return [d];
    })
    .enter().append("rect")
    .attr("class", "rect")
    .attr("width", x_values.rangeBand())
    .attr("x", function(d) {
      return 0;
    })
    .attr("y", function(d) {
      return y(d.value);
    })
    .attr("height", function(d) {
      return height - y(d.value);
    })
    .style("fill", function(d) {
      return color(d.key);
    });



});