(ns advent.day-4
  (:require
    [clojure.java.io :as io]
    [clojure.string :as s]
    [clj-time.core :as t]
    [clj-time.format :as f]
    [advent.answers :as answers]
    [advent.utils :as utils]
    [clojure.pprint :refer [pprint]]))

(def data (utils/get-input-lines "day_4"))

(defn extract-date [s]
  (let [[_ y m d h m r] (re-matches #"\[(\d+)-(\d+)-(\d+) (\d+):(\d+)] (.*)" s)]
    [(read-string (str y m d h m)) s m r]))

(defn get-shift-id [item]
  (last (re-find #"Guard #(\d+)" (:command item))))

(->> (sort data)
     (reduce
       (fn [memo line]
         (if-let [id (last (re-find #"Guard #(\d+)" line))]
           (assoc memo :curr id)
           (let [minute (->> line (re-find #":(\d+)") last Integer/parseInt)]
             (update memo :lines conj [(:curr memo) minute]))))
       {:curr nil :lines []})
     :lines
     (group-by first)
     ;; (map
     ;;   (fn [v]
     ;;     (->> v
     ;;       second
     ;;       (map second)
     ;;       frequencies
     ;;       seq
     ;;       (apply max-key second)
     ;;       )
     ;;     ))
     (map
       (fn [[id data]]
         (let [totals (->> data
                        (map last)
                        (partition 2)
                        (reduce (fn [m [a b]] (+ m (- b a))) 0))
               overlaps (->> data
                          (map second)
                          (frequencies)
                          seq
                          (apply max-key second)
                          ;; (filter (fn [[k v]] (> v 1)))
                          )
               ]
           [id totals overlaps]))
       )
     (sort-by second)
     last
     ;; ((fn [[id total overlaps]]
     ;;    (println id total (ffirst (reverse (sort-by second overlaps))))
     ;;    (* (read-string id) (ffirst (reverse (sort-by second overlaps))))
     ;;    ;; (* (read-string id) total)
     ;;    ))
     )


(defn split-command [string]
  (rest (re-find #"\[(.*?)\] (.*)" string)))

(def custom-formatter (f/formatter "yyyy-MM-dd HH:mm"))

(def sorted-line-items
  (->>
    (s/split (utils/get-input "day_4") #"\n")
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
                      :minutes minutes
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
