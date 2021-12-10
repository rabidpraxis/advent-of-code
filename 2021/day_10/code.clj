(ns day-10.code)

(def test-data
  (clojure.string/split-lines (slurp "day_10/test.txt")))

(def final-data
  (clojure.string/split-lines (slurp "day_10/input.txt")))

(def char-map
  {"(" ")"
   "[" "]"
   "{" "}"
   "<" ">"})

(def open-chars
  (set (keys char-map)))

(defn find-illegal [line]
  (loop
    [curr '()
     line (clojure.string/split line #"")]
    (if-let [char (first line)]
      (if (contains? open-chars char)
        (recur (cons char curr) (rest line))
        (if (= char (get char-map (first curr)))
          (recur (rest curr) (rest line))
          {:corrupt char}))
      {:incomplete curr})))

(def char-score
  {")" 3 "]" 57 "}" 1197 ">" 25137})

(def part-1
  (->> (map find-illegal final-data)
       (map :corrupt)
       (remove nil?)
       (map #(get char-score %))
       (apply +)))

(def char-score2
  {")" 1 "]" 2 "}" 3 ">" 4})

(def part-2
  (->>
    (map find-illegal final-data)
    (map :incomplete)
    (remove nil?)
    (map #(map (fn [e] (get char-map e)) %))
    (map
      (fn [line]
        (reduce
          (fn [curr letter]
            (+ (* curr 5) (get char-score2 letter)))
          0
          line)
        )
      )
    sort
    ((fn [v]
       (nth v (/ (count v) 2))
       ))))
