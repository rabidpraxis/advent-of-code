(ns day-2.code
  (:require
    [clojure.string :as str]))

(def final-data
  (->>
    (slurp "day_2/input.txt")
    (str/split-lines)))

(def t-input
  ["forward 5"
   "down 5"
   "forward 8"
   "up 3"
   "down 8"
   "forward 2"])

(defn process-line [line]
  (let [[t v] (str/split line #" ")]
    [t (Integer. v)]))

(defn part-1
  [input]
  (let [data
          (map process-line input)
        {:keys [d h]}
        (reduce
          (fn [ret [op v]]
            (case op
              "forward" (update ret :h + v)
              "up" (update ret :d - v)
              "down" (update ret :d + v)))
          {:d 0 :h 0}
          data)]
    (* h d)))

(part-1 final-data)

(defn part-2
  [input]
  (let [data
          (map process-line input)
        {:keys [d h]}
          (reduce
            (fn [{:keys [a] :as ret} [op v]]
              (case op
                "forward" (-> ret
                              (update :h + v)
                              (update :d + (* a v)))
                "up" (update ret :a - v)
                "down" (update ret :a + v)))
            {:d 0 :h 0 :a 0}
            data)
        ]
    (* h d)))

(part-2 final-data)
