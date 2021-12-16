(ns day-14.code
  (:require
    [clojure.string :as str]))

(def test-data
  (str/split-lines (slurp "day_14/test.txt")))

(def final-data
  (str/split-lines (slurp "day_14/input.txt")))

(defn rules [data]
  (->> (nthrest data 2)
       (map #(let [[k v] (str/split % #" -> ")]
               [(seq (char-array k)) v]))
       (into {})))

(defn grow [rules s]
  (->> (partition 2 1 s)
       (map #(if-let [in (rules %)]
               (list (first %) in)
               %))
       flatten
       (#(apply str (concat % (list (last s)))))))

(defn first-take [data i]
  (let [freqs (-> (iterate (partial grow (rules data)) (first data))
                  (nth i)
                  (frequencies)
                  (->> (sort-by second)))]
    (- (second (last freqs)) (second (first freqs)))))

(def part-1
  (time (first-take test-data 10)))

(defn new-rules [data]
  (->> (nthrest data 2)
       (map #(let [[m v] (str/split % #" -> ")]
               [m [(str (first m) v) (str v (last m))]]))
       (into {})))

(defn get-freqy [s]
  (frequencies (map #(apply str %) (partition 2 1 s))))

(defn add-or-set [v n]
  (if n (+ n v) v))

(defn step-brothers [rules freqs]
  (reduce
    (fn [coll [k n]]
      (let [[a b] (rules k)]
        (-> (update coll a (partial add-or-set n))
            (update b (partial add-or-set n)))))
    {}
    freqs))

(defn this-is-all-james [data i]
  (-> (iterate (partial step-brothers (new-rules data)) (get-freqy (first data)))
      (nth i)
      (->> (reduce
             (fn [coll [[a b] k]]
               (-> (update coll a (partial add-or-set k))
                   (update b (partial add-or-set k))))
             {})
           (map
             (fn [[l v]]
               [(bigint (Math/ceil (/ v 2))) l]))
           sort
           (#(- (first (last %)) (ffirst %))))))

(def part-2
  (this-is-all-james final-data 40))
