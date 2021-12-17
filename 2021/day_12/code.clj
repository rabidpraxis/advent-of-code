(ns day-12.code
  (:require
    [clojure.data.finger-tree :refer [double-list]]
    [clojure.string :as str]))

(def dl
  double-list)

(def test-data-1
  (str/split-lines (slurp "day_12/test.txt")))

(def test-data-2
  (str/split-lines (slurp "day_12/test_2.txt")))

(def test-data-3
  (str/split-lines (slurp "day_12/test_3.txt")))

(def final-data
  (str/split-lines (slurp "day_12/input.txt")))

(defn build-tree [data]
  (let [frags (map #(str/split % #"-") data)
        frags (concat frags (map reverse frags))]
    (->> frags
         (group-by first)
         (map
           (fn [[k v]]
             [k (apply dl (mapv second v))]))
         (into {}))))

(defn remove-end
  "Pop the last item from nested double-list, returning the new list.
  If it's the last item, pop the last item from the previous list"
  [coll]
  (let [remain (pop (last coll))]
    (if (= remain (dl))
      (remove-end (pop coll))
      (conj (pop coll) remain))))

(defn is-lower? [s]
  (= s (str/lower-case s)))

(defn part-1-ignore
  "Just ignore if next letter is in the lowercase set"
  [paths last-p]
  (-> (filter is-lower? (map last paths))
       set
       (contains? last-p)))

(defn part-2-ignore
  "Do what part 2 needs ;|"
  [paths last-p]
  (cond
    (= last-p "start")
      true
    (not (is-lower? last-p))
      false
    :else
      (let [cts (->> (conj (mapv last paths) last-p)
                     (filter is-lower?)
                     (remove #(= "start" %))
                     frequencies
                     vals)]
        (or (> (apply max cts) 2)
            (> (count (filter #(= 2 %) cts)) 1)))))

(defn find-paths [ignore-fn tree]
  (loop
    [paths (dl (dl "start") (tree "start"))
     final []]
    (let [last-p (last (last paths))]
      (cond
        (= paths (dl nil))
          final
        (= last-p "end")
          (recur (remove-end paths) (conj final (map last paths)))
        (ignore-fn (pop paths) last-p)
          (recur (remove-end paths) final)
        :else
          (recur (conj paths (tree last-p)) final)))))

(def part-1
  (->> (build-tree final-data)
       (find-paths part-1-ignore)
       count))

(def part-2
  (->> (build-tree final-data)
       (find-paths part-2-ignore)
       count))
