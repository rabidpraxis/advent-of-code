(ns day-3.code
  (:require
    [clojure.string :as str]))

(def final-data
  (->>
    (slurp "day_3/input.txt")
    (str/split-lines)))

(def test-input
  ["00100"
   "11110"
   "10110"
   "10111"
   "10101"
   "01111"
   "00111"
   "11100"
   "10000"
   "11001"
   "00010"
   "01010"])

(defn freq-col [input i]
  (frequencies (map #(nth % i) input)))

(defn all-freqs [input]
  (map-indexed (fn [i _] (freq-col input i)) (first input)))


(defn epsilon [input]
  (str/join (map
    #(->> (seq %)
          (sort-by second)
          ffirst)
    (all-freqs input))))

(defn gamma [input]
  (str/join (map
    #(->> (seq %)
          (sort-by second)
          reverse
          ffirst)
    (all-freqs input))))

(defn binary-to-int [string]
  (Integer/parseInt string 2))

(defn part-1 []
  (* (binary-to-int (epsilon final-data))
     (binary-to-int (gamma final-data))))

(defn most-common [freq]
  (if (apply = (vals freq))
    \1
    (ffirst (reverse (sort-by second (seq freq))))))

(defn least-common [freq]
  (if (apply = (vals freq))
    \0
    (ffirst (sort-by second (seq freq)))))

(defn bit-criteria [input bit-fn]
  (-> (reduce
        (fn [inputs idx]
          (if (= (count inputs) 1)
            inputs
            (let [match (bit-fn (freq-col inputs idx))
                  matches (filter
                            #(= (nth % idx) match)
                            inputs)]
              matches)))
        input
        (range (count (first input))))
      first))

(defn part-2 []
  (* (binary-to-int (bit-criteria final-data least-common))
     (binary-to-int (bit-criteria final-data most-common))))
