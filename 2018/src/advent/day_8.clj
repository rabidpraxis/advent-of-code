(ns advent.day-8
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils]))

;; (def data (utils/get-input "day_8"))
(def data "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")

(def licence (map read-string (string/split data #" ")))

(defn build-nodes [[children-ct metadata-ct & remaining]]
  (let [[new-l children] (reduce
                           (fn [[rem meta] _]
                             (let [[ll rr] (build-nodes rem)]
                               [ll (conj meta rr)]))
                           [remaining []]
                           (range children-ct))
        my-meta (take metadata-ct new-l)]
    [(drop metadata-ct new-l) {:children children
                               :metadata my-meta}]))

(defn part-1 []
  (->> (second (build-nodes licence))
       (tree-seq identity :children)
       (mapcat :metadata)
       (reduce +)))

(defn get-root-value [{:as d :keys [children metadata]}]
  (cond
    (nil? d)       0
    (seq children) (reduce + (map #(get-root-value (get children (dec %))) metadata))
    :else          (reduce + metadata)))

(defn part-2 []
  (get-root-value (second (build-nodes licence))))
