(ns day-4.code
  (:require
    [clojure.string :as str]))

(defn parse-scores [scores]
  (map #(Integer/parseInt %) (str/split scores #",")))

(defn index [item]
  (map vector (range (count item)) item))

(defn build-board [board]
  (reduce
    (fn [coll [x row]]
      (reduce
        (fn [coll [y id]]
          (assoc coll (Integer/parseInt id) [x y false]))
        coll
        (index row)))
    {}
    (index board)))

(defn split-board [lines]
  (map #(-> % (str/trim) (str/split #"\s+")) lines))

(defn boards [lines]
  (->> lines
       (partition 5 6)
       (map (comp build-board split-board))))

(defn winning-board? [board]
  (or (->> (partition 5 (sort-by first (vals board)))
           (map #(map last %))
           (map #(every? true? %))
           (some true?))
      (->> (partition 5 (sort-by second (vals board)))
           (map #(map last %))
           (map #(every? true? %))
           (some true?))))

(defn apply-score [board score]
  (if (contains? board score)
    (update board score
      (fn [[x y _]]
        [x y true]))
    board))

(defn find-first-winner [scores boards]
  (let [score (first scores)
        updated-boards (map #(apply-score % score) boards)]
    (if-let [winner (first (filter winning-board? updated-boards))]
      [score winner]
      (recur (rest scores) updated-boards))))

(defn find-last-winner [scores boards losers]
  (let [score (first scores)
        updated-boards (map #(apply-score % score) boards)]
    (if (every? winning-board? updated-boards)
      [score (apply-score (last losers) score)]
      (recur (rest scores) updated-boards (remove winning-board? updated-boards)))))

(defn unmarked-score [score board]
  (let [unmarked (filter (comp false? last last) (seq board))
        unmarked-tot (reduce #(+ %1 (first %2)) 0 unmarked)]
    (* score unmarked-tot)))

(def final-data
  (str/split-lines (slurp "day_4/input.txt")))

(def final-scores
  (parse-scores (first final-data)))

(def final-boards
  (boards (nthrest final-data 2)))

(def part-1
  (apply unmarked-score (find-first-winner final-scores final-boards)))

(def part-2
  (apply unmarked-score (find-last-winner final-scores final-boards nil)))

(comment
  (def test-data
    (str/split-lines (slurp "day_4/test.txt")))

  (def test-scores
    (parse-scores (first test-data)))

  (def test-boards
    (boards (nthrest test-data 2)))

  (apply unmarked-score (find-first-winner test-scores test-boards))
  (apply unmarked-score (find-last-winner test-scores test-boards nil)))
