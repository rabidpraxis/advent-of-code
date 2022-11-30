(ns advent.day-12
  (:require
    [clojure.string :as string]
    [advent.answers :as answers]
    [advent.utils :as utils])
  (:import java.lang.StringBuffer))

;; (set! *warn-on-reflection* true)

(def data (utils/get-input-lines "day_12"))

(defn buffer [n]
  (string/join "" (take n (repeat "."))))

(def init
  (str (buffer 100)
       (second (re-find #"initial state: (.*)" (first data)))
       (buffer 100)))

(defn init-sb [] (StringBuffer. ^String init))

(def notes (drop 2 data))

(defn ender [s]
  (subs s (dec (count s)) (count s)))

(def extract-notes
  (partial map #(subs % 0 5)))

(def add-notes (extract-notes (filter (comp (partial = "#") ender) notes)))
(def rem-notes (extract-notes (filter (comp (partial = ".") ender) notes)))

(defn replace-in-string [^StringBuffer sb ^Integer idx ^String s]
  (.replace sb ^Integer idx ^Integer (inc idx) ^String s))

(defn find-idxes [notes ^StringBuffer s ^Integer add]
  (loop [notes notes
         idxes []
         curr-idx nil]
    (if-let [note (first notes)]
      (let [new-idx (.indexOf s note (or (and curr-idx (inc curr-idx)) 0))
            idxes   (if (> new-idx 0) (conj idxes (+ new-idx add)) idxes)
            notes   (if (> new-idx 0) notes (rest notes))]
        (recur notes idxes new-idx))
      idxes)))

(defn inject [^String v]
  (fn [^StringBuffer m ^Integer idx]
    (replace-in-string m idx v)))

(def add-to-pot (inject "#"))
(def remove-from-pot (inject "."))

(defn solve [^StringBuffer sb]
  (let [idx-add (find-idxes add-notes sb 2)
        idx-rem (find-idxes rem-notes sb 2)]
    (doseq [to-rem idx-rem] (remove-from-pot sb to-rem))
    (doseq [to-add idx-add] (add-to-pot sb to-add))))

(defn nth-solve [^StringBuffer sb ^Integer ct]
  (dotimes [_ ct] (solve sb)))

(defn part-1 []
  (let [sb (init-sb)]
    (nth-solve sb 20)
    (reduce + (map #(- % 100) (find-idxes ["#"] sb 0)))))

(defn -main []
  (let [sb (init-sb)]
    (nth-solve sb 1000000000)
    (reduce + (map #(- % 100) (find-idxes ["#"] sb 0)))))
