(require
  '[clojure.string :as str])

(def test-nums
  [199 200 208 210 200 207 240 269 260 263])

(def part-1-data
  (->>
    (slurp "input.txt")
    (str/split-lines)
    (map read-string)))

(defn part-1 [nums]
  (reduce
    (fn [{:keys [before count]} i]
      (if (> i before)
        {:before i
         :count (inc count)}
        {:before i
         :count count}))
    {:before (first nums)
     :count 0 }
    (rest nums)))

(part-1 part-1-data)

(defn rollup-part-2 [data]
  (map #(reduce + %) (partition 3 1 data)))

(part-1 (rollup-part-2 part-1-data))
