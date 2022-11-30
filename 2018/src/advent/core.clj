(ns advent.core
  (:require
    [clojure.java.io :as io]
    [clojure.string :as s]
    [clj-time.core :as t]
    [clj-time.format :as f]
    [clojure.pprint :refer [pprint]]))

(defn get-input [name]
  (slurp (io/resource (str name ".txt"))))

(defn split-command [string]
  (rest (re-find #"\[(.*?)\] (.*)" string)))

(def custom-formatter (f/formatter "yyyy-MM-dd HH:mm"))

(def sorted-line-items
  (->>
    (s/split (get-input "day_4") #"\n")
    (map (fn [line]
           (let [[ts command] (split-command line)]
             {:line line
              :ts ts
              :command command
              :time (f/parse custom-formatter ts)})))
    (sort-by :time)))

(defn get-shift-id [item]
  (last (re-find #"Guard #(\d+)" (:command item))))

(def grouped-sleep-events
  (->> (reduce
         (fn [memo item]
           (let [shift-id (or (get-shift-id item) (:shift-id memo))
                 results  (or (:results memo) {})]
             (assoc memo
                    :shift-id shift-id
                    :results (update results shift-id conj item))))
         {}
         sorted-line-items)
    :results
    (map (fn [[k res]]
           [k (reverse res)]))
    (into {})))

(defn get-split-sleep-events [events]
  (partition 2 (filter #(not (get-shift-id %)) events)))

(def results
  (->> grouped-sleep-events
    (reduce
      (fn [memo [k events]]
        (let [sleep-events (get-split-sleep-events events)
              totals (map
                       (fn [[one two]]
                         (t/in-minutes (t/interval (:time one) (:time two))))
                       sleep-events)
              minutes (map
                        (fn [[one two]]
                          (range (t/minute (:time one)) (t/minute (:time two))))
                        sleep-events)
              sleep-overlap (->> (reduce
                                   (fn [memo items]
                                     (reduce
                                       (fn [memo2 item]
                                         (update memo2 item (fn [e] (if e (inc e) 0))))
                                       memo
                                       items)
                                     )
                                   {}
                                   minutes)
                              (sort-by second)
                              reverse
                              first)]
          (conj memo {:id k
                      :total (reduce + 0 totals)
                      :max-sleep-overlap sleep-overlap})))
      [])))


(println
  (->> results
    (sort-by (fn [e] (-> e :max-sleep-overlap second)))
    reverse
    first))

(println
  (->> results
    (sort-by :total)
    reverse
    first))
