(ns advent.day-5
  (:require [clojure.string :as s]
            [advent.utils :as utils]))

(def data
  (s/trim (utils/get-input "day_5")))

(defn same-letters? [pair]
  (= 1 (count (set (map s/lower-case pair)))))

(defn different-casing? [pair]
  (= 2 (count (set pair))))

(defn is-reaction? [pair]
  (and (same-letters? pair) (different-casing? pair)))

(defn run [d]
  (loop [curr (seq (s/split d #""))
         final []]
    (let [pair (take 2 curr)
          end  (nthrest curr 2)]
      (if (< (count pair) 2)
        (s/join "" (concat final pair))
        (if (is-reaction? pair)
          (recur end final)
          (recur (conj end (second pair))
                 (conj final (first pair))))))))

(defn process [string]
  (loop [t string
         pass-ct 0]
    (let [passed (run t)]
      (if (= passed t)
        passed
        (recur passed (inc pass-ct))))))


;; Part 2
(def all-letters
  (vec
    (zipmap
      (map char (range 65 91))
      (map char (range 97 123)))))

(defn drop-letters [string letter-set]
  (reduce
    (fn [st c]
      (s/replace st (str c) ""))
    string
    letter-set))

(defn get-part-2 []
  "this takes a while"
  (for [letter-set all-letters]
    [letter-set (count (process (drop-letters data letter-set)))]))
