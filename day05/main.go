package main

import (
	c "github.com/jerpa/adventofcode2024/common"
	"github.com/jerpa/adventofcode2024/day05/part1"
	"github.com/jerpa/adventofcode2024/day05/part2"
	"time"
)

func main() {
	start := time.Now()

	c.Print("Part1: ", part1.Part1(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2.Part2(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
}

var Data = `
45|19
55|17
55|38
58|41
58|46
58|72
34|56
34|63
34|83
34|33
28|12
28|56
28|61
28|41
28|38
19|85
19|22
19|97
19|28
19|46
19|62
97|18
97|81
97|34
97|96
97|58
97|55
97|79
47|76
47|95
47|81
47|41
47|83
47|79
47|96
47|63
46|93
46|34
46|18
46|84
46|83
46|45
46|64
46|29
46|26
13|41
13|96
13|47
13|84
13|86
13|46
13|76
13|85
13|52
13|79
95|56
95|48
95|28
95|22
95|77
95|44
95|13
95|11
95|62
95|26
95|86
64|41
64|34
64|45
64|17
64|16
64|32
64|18
64|89
64|26
64|93
64|84
64|51
38|47
38|58
38|12
38|81
38|85
38|49
38|96
38|76
38|61
38|64
38|48
38|13
38|72
62|58
62|55
62|48
62|52
62|76
62|72
62|46
62|89
62|12
62|31
62|97
62|29
62|86
62|81
31|63
31|54
31|17
31|28
31|33
31|83
31|11
31|42
31|34
31|38
31|32
31|27
31|44
31|95
31|49
83|13
83|56
83|51
83|61
83|32
83|11
83|27
83|33
83|26
83|54
83|77
83|97
83|28
83|19
83|62
83|48
85|51
85|41
85|11
85|32
85|16
85|52
85|84
85|93
85|45
85|34
85|95
85|77
85|17
85|83
85|44
85|55
85|31
26|97
26|12
26|56
26|54
26|22
26|33
26|49
26|86
26|76
26|28
26|47
26|62
26|27
26|81
26|79
26|58
26|19
26|48
32|26
32|61
32|48
32|81
32|42
32|17
32|28
32|44
32|62
32|13
32|27
32|79
32|56
32|54
32|19
32|97
32|49
32|11
32|47
76|18
76|93
76|46
76|63
76|45
76|84
76|72
76|29
76|34
76|96
76|64
76|95
76|17
76|51
76|41
76|83
76|85
76|55
76|16
76|31
61|79
61|52
61|97
61|31
61|47
61|22
61|64
61|86
61|29
61|72
61|96
61|55
61|48
61|85
61|46
61|81
61|41
61|45
61|76
61|13
61|58
52|11
52|28
52|95
52|18
52|55
52|31
52|16
52|32
52|29
52|93
52|38
52|83
52|19
52|26
52|34
52|17
52|51
52|27
52|84
52|44
52|63
52|77
11|97
11|28
11|33
11|72
11|46
11|96
11|49
11|58
11|61
11|64
11|22
11|79
11|13
11|76
11|56
11|86
11|38
11|47
11|48
11|62
11|42
11|12
11|81
27|58
27|19
27|46
27|33
27|64
27|81
27|97
27|56
27|22
27|11
27|28
27|61
27|62
27|42
27|12
27|76
27|86
27|44
27|79
27|13
27|48
27|38
27|47
27|49
89|28
89|32
89|93
89|16
89|11
89|34
89|29
89|52
89|84
89|54
89|95
89|63
89|55
89|18
89|44
89|27
89|17
89|77
89|42
89|83
89|31
89|51
89|26
89|19
63|61
63|17
63|11
63|44
63|16
63|26
63|27
63|51
63|77
63|13
63|54
63|28
63|42
63|48
63|95
63|19
63|62
63|49
63|32
63|38
63|56
63|33
63|97
63|83
72|18
72|95
72|31
72|85
72|32
72|17
72|27
72|41
72|51
72|93
72|63
72|26
72|16
72|89
72|29
72|52
72|83
72|55
72|44
72|77
72|45
72|84
72|54
72|34
44|97
44|76
44|62
44|38
44|48
44|49
44|46
44|79
44|13
44|86
44|19
44|64
44|12
44|11
44|56
44|61
44|33
44|28
44|58
44|47
44|42
44|96
44|81
44|22
51|22
51|42
51|79
51|97
51|13
51|26
51|27
51|48
51|54
51|56
51|47
51|28
51|44
51|38
51|61
51|62
51|12
51|17
51|19
51|86
51|33
51|49
51|11
51|32
33|48
33|13
33|46
33|89
33|81
33|55
33|96
33|76
33|47
33|97
33|52
33|22
33|56
33|61
33|79
33|86
33|64
33|85
33|45
33|41
33|12
33|72
33|58
33|62
81|55
81|77
81|58
81|89
81|84
81|83
81|76
81|41
81|85
81|46
81|63
81|18
81|45
81|51
81|31
81|64
81|95
81|93
81|34
81|72
81|96
81|52
81|16
81|29
86|84
86|83
86|47
86|81
86|31
86|55
86|76
86|93
86|85
86|16
86|45
86|58
86|41
86|79
86|34
86|72
86|18
86|63
86|96
86|46
86|64
86|52
86|89
86|29
41|45
41|55
41|17
41|89
41|77
41|19
41|31
41|51
41|11
41|83
41|95
41|63
41|44
41|29
41|26
41|18
41|54
41|34
41|16
41|93
41|32
41|27
41|84
41|52
77|48
77|22
77|62
77|17
77|12
77|44
77|61
77|19
77|49
77|27
77|54
77|86
77|47
77|32
77|42
77|28
77|38
77|56
77|33
77|11
77|13
77|51
77|26
77|97
84|11
84|28
84|49
84|27
84|93
84|83
84|32
84|61
84|42
84|33
84|95
84|77
84|17
84|56
84|44
84|54
84|38
84|18
84|19
84|26
84|51
84|34
84|16
84|63
54|97
54|19
54|79
54|62
54|49
54|46
54|48
54|11
54|61
54|22
54|47
54|42
54|58
54|13
54|38
54|44
54|81
54|56
54|33
54|86
54|12
54|27
54|28
54|76
18|93
18|95
18|63
18|54
18|17
18|19
18|28
18|77
18|61
18|11
18|51
18|56
18|42
18|26
18|83
18|13
18|27
18|33
18|44
18|62
18|32
18|16
18|38
18|49
12|18
12|84
12|89
12|81
12|63
12|86
12|41
12|85
12|58
12|22
12|55
12|93
12|29
12|46
12|34
12|64
12|52
12|96
12|79
12|72
12|76
12|31
12|47
12|45
93|77
93|32
93|83
93|42
93|63
93|13
93|61
93|17
93|26
93|49
93|11
93|28
93|62
93|54
93|16
93|48
93|38
93|56
93|19
93|33
93|44
93|51
93|95
93|27
56|46
56|48
56|86
56|41
56|13
56|76
56|79
56|96
56|52
56|85
56|62
56|47
56|55
56|31
56|89
56|61
56|12
56|58
56|64
56|97
56|22
56|81
56|72
56|45
96|26
96|45
96|27
96|89
96|55
96|51
96|31
96|34
96|41
96|54
96|85
96|63
96|95
96|29
96|84
96|83
96|18
96|77
96|72
96|17
96|52
96|16
96|32
96|93
48|18
48|12
48|22
48|85
48|58
48|79
48|46
48|41
48|29
48|96
48|89
48|47
48|72
48|31
48|34
48|97
48|55
48|84
48|45
48|86
48|76
48|64
48|52
48|81
16|61
16|17
16|28
16|32
16|95
16|44
16|48
16|19
16|33
16|77
16|54
16|51
16|56
16|13
16|26
16|83
16|12
16|62
16|97
16|42
16|38
16|27
16|49
16|11
79|58
79|16
79|96
79|84
79|85
79|31
79|83
79|55
79|34
79|64
79|77
79|52
79|72
79|41
79|29
79|95
79|45
79|89
79|46
79|18
79|76
79|93
79|81
79|63
49|96
49|86
49|72
49|46
49|52
49|85
49|13
49|56
49|41
49|62
49|48
49|79
49|47
49|12
49|45
49|33
49|81
49|22
49|97
49|64
49|58
49|76
49|89
49|61
17|26
17|54
17|61
17|62
17|22
17|58
17|12
17|47
17|33
17|11
17|19
17|48
17|49
17|97
17|13
17|86
17|56
17|38
17|42
17|79
17|28
17|27
17|44
17|81
29|93
29|27
29|32
29|54
29|11
29|63
29|44
29|49
29|17
29|42
29|18
29|77
29|34
29|38
29|83
29|19
29|28
29|33
29|16
29|51
29|26
29|95
29|56
29|84
42|13
42|62
42|22
42|76
42|64
42|38
42|72
42|47
42|12
42|97
42|48
42|79
42|28
42|86
42|46
42|33
42|61
42|96
42|81
42|41
42|58
42|56
42|49
42|85
22|34
22|89
22|18
22|76
22|86
22|81
22|63
22|45
22|84
22|64
22|79
22|55
22|16
22|41
22|96
22|93
22|52
22|58
22|47
22|29
22|72
22|85
22|31
22|46
45|52
45|55
45|54
45|11
45|44
45|17
45|77
45|27
45|26
45|16
45|93
45|42
45|31
45|51
45|84
45|95
45|18
45|32
45|63
45|89
45|29
45|34
45|83
55|63
55|16
55|27
55|31
55|51
55|54
55|83
55|84
55|42
55|18
55|28
55|44
55|95
55|93
55|19
55|49
55|26
55|11
55|32
55|29
55|34
55|77
58|95
58|76
58|31
58|93
58|96
58|77
58|83
58|52
58|45
58|18
58|34
58|32
58|29
58|89
58|84
58|51
58|85
58|64
58|55
58|16
58|63
34|95
34|16
34|17
34|19
34|77
34|28
34|54
34|27
34|26
34|11
34|38
34|93
34|44
34|32
34|42
34|49
34|62
34|51
34|18
34|61
28|62
28|76
28|85
28|81
28|22
28|13
28|47
28|97
28|86
28|33
28|58
28|46
28|96
28|79
28|45
28|72
28|49
28|48
28|64
19|12
19|64
19|58
19|79
19|76
19|49
19|81
19|38
19|48
19|33
19|61
19|47
19|42
19|96
19|56
19|86
19|13
19|72
97|22
97|47
97|45
97|93
97|29
97|84
97|46
97|52
97|76
97|89
97|85
97|86
97|31
97|12
97|72
97|41
97|64
47|52
47|29
47|58
47|45
47|46
47|84
47|31
47|72
47|89
47|18
47|93
47|34
47|64
47|16
47|85
47|55
46|85
46|96
46|17
46|41
46|77
46|63
46|32
46|16
46|89
46|31
46|52
46|72
46|51
46|95
46|55
13|64
13|97
13|12
13|29
13|45
13|55
13|31
13|89
13|72
13|81
13|48
13|22
13|34
13|58
95|49
95|19
95|54
95|61
95|17
95|97
95|51
95|12
95|32
95|27
95|38
95|42
95|33
64|54
64|31
64|52
64|55
64|72
64|85
64|77
64|95
64|96
64|63
64|29
64|83
38|89
38|62
38|45
38|22
38|41
38|79
38|46
38|33
38|97
38|56
38|86
62|22
62|41
62|85
62|96
62|64
62|13
62|45
62|79
62|84
62|47
31|29
31|19
31|16
31|77
31|51
31|84
31|93
31|18
31|26
83|17
83|42
83|44
83|38
83|95
83|12
83|49
83|22
85|63
85|54
85|89
85|26
85|27
85|29
85|18
26|44
26|42
26|11
26|38
26|13
26|61
32|33
32|22
32|86
32|38
32|12
76|89
76|32
76|52
76|77
61|89
61|12
61|62
52|42
52|54
11|19

56,79,55,52,85,41,61,97,64,72,86,46,58,48,96,62,76,12,13
16,95,77,51,17,49,33,56,97
95,38,49,61,32,19,77,22,13,27,56
84,22,89,64,45,79,34,85,72,96,48,55,29
84,34,18,93,63,16,95,77,51,32,17,26,54,27,44,11,19,42,28,38,49,33,56
79,72,96,45,76,58,64,46,55,84,86,63,22,52,18,93,85,31,47,41,29,89,81
33,34,84,32,49,93,54,63,95,27,83,17,77,28,29,38,42,26,44,11,18
34,18,93,63,83,95,51,32,26,27,11,19,42,28,49,33,56
31,29,84,18,93,63,16,95,51,26,54,27,19,42,28,38,49
31,55,13,22,84,47,72
95,77,51,32,17,26,54,27,44,11,42,28,38,49,33,56,61,62,13,48,97,12,22
34,18,93,63,16,77,51,32,17,26,54,44,11,28,38,49,33,56,61
45,89,84,93,77
95,29,84,85,16,83,96,41,89,64,31,52,77,76,34,72,32
51,32,17,26,54,27,44,11,19,42,28,38,49,33,56,61,13,48,97,12,22
13,77,38,49,62,95,12,32,48,54,44,56,17,11,19,28,22,26,51,42,61,97,27
41,45,89,52,55,31,29,84,34,18,93,83,95,77,51,17,26,54,27,44,11
85,89,31,29,84,34,63,16,95,77,51,27,44
72,62,47,49,97,45,76,61,79,38,13,22,56
89,18,46,55,63,72,85,41,64,83,16,52,45,51,96,17,31,93,29
58,41,45,52,55,18,93,83,95,77,51
97,12,22,86,47,79,81,58,76,46,64,96,72,85,41,45,89,52,55,31,29,84,18
79,81,58,76,46,64,72,85,41,45,89,52,55,31,29,84,34,93,63,16,83
44,58,28,54,86,79,27,22,49,62,47,19,11,38,13,42,48,26,61,81,33
19,64,49,56,28,11,46,44,22,33,62,47,48,79,86
42,26,28,52,83,95,29,44,54
11,33,54,32,27,51,16,19,42,95,61,44,56,97,77,13,28,38,49
19,31,27,95,55,44,77,93,54,63,26,45,52,29,17,83,34,84,18,11,16,32,89
18,63,16,77,51,32,54,27,44,42,49,33,56,61,62
49,58,72,86,76,46,13,12,33,47,38,28,79,56,48,97,42,62,61,64,85
33,56,61,13,48,97,12,22,79,58,76,64,41,45,52
77,26,51,85,41,54,27,17,34,32,44
19,62,33,54,95,97,13,26,12,77,32,83,56
17,26,27,44,19,42,28,38,49,33,56,62,13,48,97,22,86,47,81
72,48,89,29,97,86,47,22,76,79,58,85,81,41,12,84,31,96,55,45,34,52,64
77,17,27,11,19,42,33
95,84,93,27,83,72,34,16,45,55,31,54,52,89,32
84,63,77,83,38,42,32,27,16,19,26,11,49,34,44,33,17,95,56
26,54,27,44,19,42,33,56,13,48,97,12,22,86,47,79,81
47,85,34,18,79,72,64,76,97,46,96
48,12,17,86,47,62,33,38,32,54,49,11,97,79,44
28,49,33,56,61,62,13,22,47,76,46,72,41
85,41,45,89,52,55,31,29,84,34,18,93,63,16,83,95,51,32,17,26,54,27,44
81,46,64,45,52,55,29,34,16,95,77
27,93,95,18,54,55,83,34,28,52,63,29,31,77,26
12,55,56,13,89,47,85,46,48,64,58,45,79
84,16,85,31,52
47,52,97,31,64,29,41,76,79,46,48,96,86,34,12
32,93,63,49,44,31,29,42,54
31,84,34,18,63,16,83,95,77,51,32,17,54,27,44,11,19,42,28,38,49
41,52,55,31,29,84,34,18,93,63,16,95,32,26,27,44,11
41,29,34,32,17,54,11
42,38,49,56,61,62,13,48,97,86,47,79,81,58,76,46,96,72,85
89,52,31,93,16,17,26,54,19
77,42,28,49,97,22,86
85,89,22,76,47,34,84,81,55,63,29,52,72,18,79,96,58,86,45,31,41
28,38,49,33,56,61,62,13,48,97,12,86,81,58,76,46,64,96,72,85,41
32,93,41,52,89,77,64,31,83,84,29,18,63,85,72,34,96,46,45,17,55
77,51,26,33,48,97,86
18,63,32,17,19,38,33,56,62
51,32,29,77,31,54,83,93,84,27,18
93,63,83,95,32,27,28,49,61
38,49,33,56,62,13,48,97,12,86,47,79,58,76,46,64,96,72,85,41,45
17,63,77,61,42,19,38,27,49,33,11,32,16,83,93,51,62,54,56
28,81,12,33,19,97,44,79,42,47,46,11,22,49,38,64,61,62,56,48,76,13,86
13,49,72,19,22,76,42,97,61,12,48,81,62,86,96,28,33,46,47,79,58,64,38
64,12,72,96,76,46,45
27,11,56,62,12,86,47
84,34,18,93,63,16,83,95,77,51,32,17,26,54,27,44,11,19,28,38,49,33,56
96,34,52,45,83,16,26,93,64,55,95,29,72,31,51
62,72,89,96,76,86,64,45,22,56,79,12,41,47,85,52,97,55,61
11,19,49,56,61,22,79,81,58,64,96
58,76,46,72,85,41,89,52,55,31,29,34,18,93,16,83,95,77,51
56,61,62,13,48,22,79,81,76,96,85,41,89,52,55
26,17,96,18,54,45,16,84,72,93,95
51,32,17,27,44,11,38,49,33,56,62,13,48,12,22,86,47
63,16,83,95,77,51,32,17,26,54,27,19,42,28,49,33,61,62,13
55,29,84,63,17,54,38
13,48,97,12,22,86,47,79,81,76,46,64,96,72,85,41,45,89,52,55,31,29,84
32,55,16,51,89,18,93,63,83,52,85,77,95,84,54,96,45,26,29,72,41
54,27,44,19,42,28,38,61,62,13,97,12,22,86,79
47,96,52,55,29,34,83
12,22,47,79,46,96,72,85,41,89,84
28,51,38,56,17,49,95,26,42,11,77,97,19,44,27,83,54,32,12,61,33,48,13
64,83,34,26,77,18,55
28,33,56,61,48,97,22,86,47,79,81,58,76,46,64,96,72,85,41
38,62,22,19,27,49,28,47,13
28,83,17,16,26,93,61,34,51
26,49,62,48,97,12,22,86,58
97,61,55,52,41,62,22,45,85,64,47,96,13,56,76,81,58
63,95,77,51,32,17,26,54,27,44,11,42,28,38,49,33,56,61,62,13,48
86,56,44,62,28,13,33,17,79,54,12,49,11,22,38,32,42,61,48
11,38,42,54,27,22,62,12,61,19,48,44,49,86,56,81,17,97,33,47,13,28,26
31,18,93,63,95,77,51,17,26,54,27,11,19,38,49
63,51,31,55,64,16,84,18,83,76,72,29,85,46,32,52,34
13,97,86,79,64,72,85,41,84
55,34,77,51,32,54,11,19,42
55,95,84,51,16,77,38,27,93,19,54,34,42
62,11,56,13,38,26,97,48,54,86,27,61,22,17,49,19,77
54,77,62,27,56,17,16,83,26,97,32,61,95
45,89,52,31,29,84,18,93,16,83,95,77,51,32,17,54,27,11,19
17,26,27,11,19,28,49,56,61,62,13,48,12,22,47
96,72,85,52,55,31,29,84,63,51,17,26,54
22,86,47,79,81,58,76,46,64,96,72,85,41,45,52,55,31,29,34,18,63
55,76,79,52,85
31,29,84,34,18,93,63,16,83,95,77,51,17,26,54,27,44,11,42,28,49
42,48,13,79,47,19,56,49,11,96,86
95,77,51,17,54,44,13,48,97,12,22
41,31,84,16,83,17,54,44,11
12,41,86,52,93,46,89
51,61,42,62,18,54,49,95,83,17,56
81,58,76,46,64,96,72,85,41,45,89,52,55,31,29,34,18,93,63,16,83,95,77
51,77,42,32,54,17,83,27,49,28,95,61,97,19,48
11,54,48,51,95,27,17,62,44,28,32,19,61,83,13,49,97,42,33,77,38,26,16
18,63,16,77,51,17,27,11,19,42,38,49,33,56,62
86,47,79,81,58,76,46,64,96,72,85,45,89,52,55,31,29,84,34,18,93,63,16
95,89,41,76,16,63,31,79,52,64,46,72,45,81,55,96,83
55,84,63,83,17,44,11,42,38
64,33,28,56,38,47,97,12,86,96,61,81,22,13,79,76,62,11,46,42,48,49,58
31,97,12,81,41,46,58,89,61
97,22,81,46,64,72,29
48,42,64,56,58,11,76,96,47,49,33,81,38,19,86,46,97
77,51,32,11,49,33,56,62,13,97,86
58,29,72,79,76,93,86,84,52,89,12,34,85,18,46
72,45,89,52,55,31,29,34,18,93,63,16,83,77,17,26,27
33,19,13,77,56,26,11
89,55,72,31,29,85,12,45,41,22,76,96,86,13,62,79,52,81,47,58,46,48,97
79,64,96,41,45,29,84,63,16
44,26,27,83,95,54,49,33,34,42,29
89,48,79,47,46,22,72,86,97,33,49,96,62
81,45,46,72,93,52,58,79,95,76,63,83,89,84,31,55,41,29,34,96,85
33,86,96,81,58,85,72,97,13,46,48,79,76,61,12,22,64,89,52,47,62,45,41
76,64,72,85,41,89,52,55,34,93,16,83,95,77,32
33,56,61,48,97,12,22,86,47,79,81,76,46,89,52
34,95,51,54,27,42,33,56,61
77,51,32,17,26,54,27,44,11,19,42,28,38,49,33,56,61,62,13,48,97,12,22
63,51,55,76,58,45,46,16,95,64,72,96,34
42,13,11,54,49,27,62,38,83,26,93,77,32,28,16,51,63,95,33,17,61,19,44
63,51,38,31,54,84,29,55,28,44,77
61,62,13,48,97,22,47,81,58,46,96,72,41,89,52,55,31
97,12,22,86,47,81,58,76,46,64,72,85,41,45,89,52,31,29,84,34,18
31,93,16,95,26,27,49
48,33,56,51,26,28,38,63,62,42,54,11,44,95,77
79,42,28,96,12,86,56,47,38,49,81,33,72,48,61
81,58,76,64,41,45,55,31,84,34,93,63,83,95,77
72,13,79,58,19,47,49,12,46,38,42,76,64,56,62,81,48,28,61,96,22
84,93,63,16,95,17,19,38,56
95,84,89,45,41,76,79
33,54,97,86,79,38,27,17,22,47,62,13,44,49,28,42,48,11,32,12,26,61,56
34,18,93,63,83,95,77,32,17,26,54,27,11,19,42,49,33,56,61
89,79,81,49,85,58,12,76,45
89,95,93,84,44,18,51,31,54,63,11,29,55,26,27,83,34
62,47,81,58,76,85,45,52,55,31,29
81,85,76,79,55,97,45,41,86,62,13,96,72,89,12,29,22,47,64,31,52,58,46
38,11,27,58,49,81,97,28,48,47,86,42,79,13,61,44,76,12,56,19,62,22,54
54,61,11,95,44,49,26,62,51,18,63,56,28
13,48,97,22,86,47,79,81,58,76,46,64,96,72,85,41,45,89,52,55,31,29,84
29,41,63,44,18,11,32,45,27
85,89,32,18,83,95,51,55,72,93,64,84,63,26,96,34,41,16,45,31,77
44,47,64,13,49,33,56
79,55,52,58,46,95,41
18,45,29,84,93,89,51,32,77,34,96,63,95,85,26,52,55,64,72
77,51,32,17,26,54,27,44,11,19,28,38,49,33,61,62,48,22,86
54,28,49,33,61,62,97,22,76
44,19,27,79,42,13,54,22,49,33,47,81,86,48,38,28,62,61,56,12,97,58,11
33,56,61,62,48,22,79,58,72
38,28,48,12,26,58,86
12,42,38,22,51,47,61
56,38,62,13,86,79,27,28,44,12,97,81,46,49,33,47,61
18,76,51,29,52,83,41
13,48,97,12,86,47,81,46,96,72,85,45,84
97,12,22,86,47,79,81,58,76,46,96,72,85,45,89,52,55,31,29,34,18
58,47,19,12,26,11,79,42,13,44,22,33,38,54,49,56,61,62,86,27,48
11,19,42,28,38,49,33,61,62,13,48,97,12,22,86,47,79,81,58,76,46,64,96
28,49,56,61,13,48,86,81,58,76,46
`
