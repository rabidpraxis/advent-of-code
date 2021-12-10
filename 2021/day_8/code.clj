(ns day-8.code
  (:require
    [clojure.set :as s]
    [clojure.string :as str]))

;; (def number-segments
;;   {1 #{:c :f}
;;    7 #{:a :c :f}
;;    4 #{:b :c :d :f}
;;    8 #{:a :b :c :d :e :f :g}
;;    3 #{:a :c :d :f :g}
;;
;;    5 #{:a :b :d :f :g}
;;    9 #{:a :b :c :d :f :g}
;;    2 #{:a :c :d :e :g}
;;    6 #{:a :b :d :f :e :g}
;;    0 #{:a :b :c :e :f :g}
;;    })
;;
;; (map
;;   #(s/subset? (seg 7) %)
;;   (vals (select-keys number-segments [2 0])))
;;
;; (defn seg [i]
;;   (get number-segments i))
;;
;; (s/subset? (seg 1) (seg 5))
;; (s/difference (seg 7) (seg 1))
;; (s/difference (seg 5) (seg 3))
;;
;; {:a (s/difference (seg 7) (seg 1))
;;  :b (s/difference (seg 4) (seg 3))
;;  :c (s/difference (seg 4) (seg 8))}

;; Find 3
;; (s/subset? (seg 4) (seg 2))

(def test-data
  (clojure.string/split-lines (slurp "day_8/test.txt")))

(def final-data
  (clojure.string/split-lines (slurp "day_8/input.txt")))

(def test-line
  "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")

(def part-1
  (->> final-data
       (mapcat #(-> %
                    (clojure.string/split #"\| ")
                    last
                    (clojure.string/split #" ")
                    ))
       (map count)
       (filter #(contains? (set [2 3 4 7]) %))
       count))

(defn extract-segs
  [line]
  (-> line
      (str/split #" ")
      (->> (filter (partial not= "|"))
           (map #(set (map keyword (str/split % #"")))))
      distinct))

(def unique-counts
  {2 1, 3 7, 4 4, 7 8})

(defn extract-by-counts
  "Find the first 4 numbers by counts"
  [line]
  (let [found (group-by
                (fn [l]
                  (or (get unique-counts (count l))
                      :else))
                line)
        matched (->> (dissoc found :else)
                     (map (fn [[k v]]
                            [k (first v)]))
                     (into {})) ]
    [matched (:else found)]))

(defn find-by
  [subset-fn subset-n length match-n [matches segs]]
  (let [found (first (filter
                       #(and (subset-fn (get matches match-n) %)
                             (= (count %) length))
                       segs))
        segs (remove #(= found %) segs)]
    [(assoc matches subset-n found) segs]))

(defn segments [line]
  (->> line
       extract-segs
       extract-by-counts
       (find-by s/subset? 3 5 1)
       (find-by s/subset? 9 6 3)
       (find-by s/superset? 5 5 9)
       (find-by s/subset? 6 6 5)
       (find-by s/subset? 0 6 7)
       (find-by (constantly true) 2 5 1)
       first))

(defn find-output-value [line]
  (let [idigits (s/map-invert (segments line))
        matches (-> line
                    (str/split #"\| ")
                    second
                    (str/split #" ")
                    (->> (map #(set (map keyword (str/split % #""))))))]
    (->> (map #(get idigits %) matches)
        (str/join "")
        (#(Integer/parseInt %)))))

(def part-2
  (apply + (map find-output-value final-data)))
