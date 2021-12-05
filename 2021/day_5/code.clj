(ns day-5.code
  (:require
    [clojure.string :as str]))

(def final-data
  (str/split-lines (slurp "day_5/input.txt")))

(def test-data
  (str/split-lines (slurp "day_5/test.txt")))

(defn extract-segment
  [line]
  (->> line
       (re-find #"(\d+),(\d+) -> (\d+),(\d+)")
       rest
       (map #(Integer/parseInt %))))

(defn inclusive-range
  ([start end]
   (if (> start end)
     (range start (dec end) -1)
     (range start (inc end)))))

(defn ranges
  [with-diag? [x1 y1 x2 y2]]
  (cond
    (= x1 x2)
      (for [y (inclusive-range y1 y2)] [x1 y])
    (= y1 y2)
      (for [x (inclusive-range x1 x2)] [x y1])
    with-diag?
      (map vector (inclusive-range x1 x2) (inclusive-range y1 y2))
    :else
      []))

(defn points [data with-diag?]
  (->> data
       (map extract-segment)
       (mapcat (partial ranges with-diag?))
       frequencies
       vals
       (filter #(>= % 2))
       count))

(def part-1 (points final-data false))
(def part-2 (points final-data true))
