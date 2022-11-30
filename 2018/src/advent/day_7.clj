(ns advent.day-7
  (:require
    [clojure.string :as string]
    [clojure.set :as set]
    [com.stuartsierra.dependency :as dep]
    [advent.test-case :as test-case]
    [advent.utils :as utils]))

(def data (utils/get-input-lines "day_7"))
;; (def data test-case/day-7)

(defn get-paths [s]
  (let [[_ from to] (re-matches #"Step (.).*?step (.).*" s)]
    [from to]))

(def sets
  (reduce
    (fn [[mapped counts] item]
      (let [[from to] (get-paths item)]
        [(update mapped from conj to)
         (update counts to (fnil inc 0))]))
    [{} {}]
    data))

(def branches (first sets))
(def counts (second sets))

(def all
  (apply sorted-set (set/union (set (keys branches)) (set (keys counts)))))

(def starting
  (apply sorted-set (set/difference (set (keys branches)) (set (keys counts)))))

(defn inject-branches [root-step count-set work-set]
  (reduce
    (fn [memo step]
      (let [dec-memo (update-in memo [:cs step] dec)]
        (if (zero? (get-in dec-memo [:cs step]))
          (update dec-memo :ws conj step)
          dec-memo)))
    {:cs count-set :ws work-set}
    (sort (get branches root-step))))

(defn part-1 []
  (loop [root-set  starting
         count-set counts
         final     ""]
    (if-let [root-step (first root-set)]
      (let [{:keys [cs ws]} (inject-branches root-step count-set root-set)]
        (recur (disj ws root-step) cs (str final root-step)))
      final)))

(def range-start 61)
;; (def range-start 1)

(def step-times
  (zipmap all (range range-start 200)))

(def working?
  seq)

(defn complete-steps [workers t]
  (filter
    (fn [[step start]]
      (= (+ start (get step-times step)) t))
    workers))

(defn clear-done-work [workers complete]
  (vec (remove (set complete) (set workers))))

(defn assign-to-workers [workers ws t]
  (let [spaces (- 5 (count workers))
        steps  (map #(do [% t]) (take spaces ws))]
    (if (seq steps)
      [(concat workers steps) (apply sorted-set (drop spaces ws))]
      [workers ws])))

;; (assign-to-workers [["A" 0]] (sorted-set "B" "C") 2)
;; (clear-done-work nil [["A" 0]])
;; (complete-steps [["B" 0] ["A" 1]] 2)

(defn part2 []
  (loop [work-set  starting
         count-set counts
         workers   []
         t         0]
    (println work-set count-set workers t)
    (if (or (first work-set) (working? workers))
      (let [complete-steps  (complete-steps workers t)
            updated-workers (clear-done-work workers complete-steps)
            {:keys [cs ws]} (reduce
                              (fn [{:keys [cs ws]} [step _]]
                                (inject-branches step cs ws))
                              {:cs count-set :ws work-set}
                              complete-steps)
            [workers ws] (assign-to-workers updated-workers ws t)]
        (recur ws cs workers (inc t)))
      (- t 1))))
